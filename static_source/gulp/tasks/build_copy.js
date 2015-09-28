/**
 * Created by delta54 on 01.12.14.
 */
var gulp = require('gulp'),
    conf = require('../config').copy;

gulp.task('build_copy', function() {
    return gulp.src(conf.source,
        {base: conf.source_dir})
        .pipe(gulp.dest(conf.dest));
});