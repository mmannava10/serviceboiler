var gulp = require('gulp');
var path = require('path');
var shell   = require('gulp-shell');
var gulp        = require('gulp');
var extend      = require('extend');
var parseArgs   = require('minimist');
var runSequence = require('run-sequence');
var gulpif      = require('gulp-if');
var uglify      = require('gulp-uglify');
var rename      = require('gulp-rename');

// 
var exec = require('child_process').exec;
var generateProtoBufs = 'protoc -I=./pb --go_out=./src/pb/db ./pb/db.proto'

// Your setup related variables
var devpath     = 'C:/Users/mmannava/github.com/mmannava10/serviceboiler'



// Configuration
//
var config = extend({
  env: process.env.NODE_ENV
}, parseArgs(process.argv.slice(2)));

// Getters / Setters
//

gulp.task( 'set-project-env', function() {
 return  process.env.GOPATH = devpath

  
});

gulp.task('generate-protobuf', function() {
  console.log("Generating protobufs...")
  console.log("Running: " + generateProtoBufs)
  exec(generateProtoBufs, function (err, stdout, stderr) {
    console.log(stdout);
    console.log(stderr);

  });
});

gulp.task('set-dev-node-env', function() {
  return process.env.NODE_ENV = config.env = 'development';
});
gulp.task('set-prod-node-env', function() {
  return process.env.NODE_ENV = config.env = 'production';
});

// General tasks
//
gulp.task('scripts', function() {
  return gulp.src('src/scripts/main.js')
     .pipe( gulpif(config.env === 'production', uglify() ))
     .pipe( gulpif(config.env === 'production', rename({suffix:'.min'}) ))
     .pipe( gulp.dest('dest/scripts') );
});

// Run tasks
//
gulp.task('build', ['scripts']);
gulp.task('develop', ['set-dev-node-env'], function() {
  return runSequence(
     'build'
  );
});     // => main.js
gulp.task('deploy', ['set-prod-node-env'], function() {
  return runSequence(
     'build'
  );
}); 

var goPath = 'src/**/**/**/*.go';

gulp.task('compilepkg', function() {
  return gulp.src(goPath, {read: false})
    .pipe(shell([ 'go install -v  <%= stripPath(file.path) %>'],
      {
          templateData: {
            stripPath: function(filePath) {
             // console.log("B:process.env.GOPATH=[" + process.env.GOPATH  + "]")
              // process.env.GOPATH = devpath
              //console.log("A:process.env.GOPATH=[" + process.env.GOPATH  + "]")
              var subPath = filePath.substring(process.cwd().length + 5);
              var pkg = subPath.substring(0, subPath.lastIndexOf(path.sep));
              console.log("GULP:pk=" + pkg)
              return pkg;
            }
          }
      })
    );
});

gulp.task('watch', function() {
  gulp.watch(goPath, ['compilepkg']);
});

gulp.task('default', ['generate-protobuf', 'compilepkg']);