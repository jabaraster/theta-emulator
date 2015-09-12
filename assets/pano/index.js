$(function () {
  sizing();
  $(window).resize(function() {
    sizing();
  });
});

function sizing(){
  $("#canvas").attr({height:$("#wrapper").height()});
  $("#canvas").attr({width:$("#wrapper").width()});
}

var timer = false;
$(window).resize(function() {
    if (timer !== false) {
        clearTimeout(timer);
    }
    timer = setTimeout(function() {
        console.log('resized');
        init_pano('canvas');
    }, 200);
});
