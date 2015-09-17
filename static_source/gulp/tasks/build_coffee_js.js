/**
 * Created by delta54 on 6/30/15.
 */

var gulp = require('gulp'),
    conf = require('../config').build_coffee_js,
    concat = require('gulp-concat'),
    ngconcat = require('gulp-ngconcat'),
    inject = require('gulp-inject');
    coffee = require('gulp-coffee');
    uglify = require('gulp-uglify'),
    ngClassify = require('gulp-ng-classify'),
    gutil = require('gulp-util');

gulp.task('build_coffee_js', function() {
    return gulp.src(conf.source)

        .pipe(coffee({bare: true})
            .on('error', gutil.log))     // Compile coffeescript
        //.pipe(ngconcat(conf.filename))
        .pipe(concat(conf.filename))
        //.pipe(uglify())
        //.pipe(ngClassify())
        .pipe(gulp.dest(conf.dest));

});
