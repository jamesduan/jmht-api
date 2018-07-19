$(function() {
    /*======右按钮======*/
    $(".slid-list-div .right").click(function() {
        nextscroll();
    });

    function nextscroll() {
        var vcon = $("#slid1");
        var offset = ($("#slid1 li").width()) * -1;
        $("#slid1 ul").stop().animate({
            marginLeft: offset
        }, "slow", function() {
            var firstItem = $("#slid1 li").first();
            vcon.find("ul").append(firstItem);
            $(this).css("margin-left", "0px");
        });
    };

    /*========左按钮=========*/
    $("slid-list-div .left").click(function() {
        var vcon = $("#slid1");
        var offset = ($("#slid1 li").width() * -1);
        var lastItem = $("#slid1 ul li").last();
        vcon.find("ul").prepend(lastItem);
        $("#slid1 ul").css("margin-left", offset);
        $("#slid1 ul").animate({
            marginLeft: "0px"
        }, "slow")
    });
});


$(function() {
    /*======右按钮======*/
    $(".slid-list-div .right1").click(function() {
        nextscroll();
    });

    function nextscroll() {
        var vcon = $("#slid2");
        var offset = ($("#slid2 li").width()) * -1;
        $("#slid2 ul").stop().animate({
            marginLeft: offset
        }, "slow", function() {
            var firstItem = $("#slid2 ul li").first();
            vcon.find("ul").append(firstItem);
            $(this).css("margin-left", "0px");
        });
    };
    /*========左按钮=========*/
    $(".slid-list-div .left1").click(function() {
        var vcon = $("#slid2");
        var offset = ($("#slid2 li").width() * -1);
        var lastItem = $("#slid2 ul li").last();
        vcon.find("ul").prepend(lastItem);
        $("#slid2 ul").css("margin-left", offset);
        $("#slid2 ul").animate({
            marginLeft: "0px"
        }, "slow")
    });
});