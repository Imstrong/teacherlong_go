$(function () {
    $(document).click(function(e){
        var obj = e.target || e.srcElement;
        for(index in obj.classList) {
            if (obj.classList[index] == "wechat") {
                return;
            }
            $(".qr-code").fadeOut();
        }
    })
    $(".wechat").click(function(){
        $(".qr-code").fadeIn();
    });
    $(".after-wrapper").click(function(){
        var photos=$(".photo-group").children("div");
        var first=$(photos[photos.length-1]).clone();
        $(photos[photos.length-1]).fadeOut("sloe",function(){
            $(".photo-group").prepend(first);
            $(photos[photos.length-1]).remove();
        })
        $(".desc:not(:last-child)").css("display","none");
    })
})