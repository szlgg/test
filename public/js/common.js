// 全局管理对象
gf = {
    // 刷新验证码
    reloadCaptcha: function() {
        //创建Ajax对象
        const xhr = new XMLHttpRequest();
        //告诉Ajax对象以什么方式向哪个地址发送请求
        xhr.open('get', "/captcha?v=="+Math.random());
        //向服务器端发送请求
        xhr.send();
        xhr.onload=function(){
            $("img.captcha").attr("src", JSON.parse(xhr.response));
        }
    },
}

// 初始化请求
$(document).ready(
    function(){
        $.post("/init", "",
            function (data) {
                if (data.isLogin === true) {
                    $(".logged-in").css("display","flex");
                    $(".user-image").after(data.User.Nickname).attr("src", data.User.Avatar);
                } else {
                    $(".not-logged").css("display","flex");
                }
            });
    }
);

// 用户模块
gf.user = {
    // 退出登录
    logout: function () {
        console.log("开始")
        layer.confirm('您确定需要注销当前登录状态吗？', {icon: 0, title:'注销登录'}, function(index){
            alert("执行")
            // window.location.href = "/user/logout";
            window.location.replace("/user/logout")
        });
    },
}

// 互动模块
gf.interact = {
    // 检查赞
    checkZan: function (elem, id) {
        var type = $(elem).attr("data-type")
        if ($(elem).find('.icon').hasClass('icon-zan')) {
            this.cancelZan(elem, type, id)
        } else {
            this.zan(elem, type, id)
        }
    },
    // 赞
    zan: function (elem, type, id) {
        jQuery.ajax({
            type: 'POST',
            url : '/interact/zan',
            data: {
                id:   id,
                type: type
            },
            sync: true,
            success: function (r, status) {
                if (r.code == 0) {
                    let number = $(elem).find('.number').html()
                    $(elem).find('.number').html(parseInt(number)+1)
                    $(elem).find('.icon').removeClass('icon-zan1').addClass('icon-zan')
                } else {
                    layer.alert(r.message, {icon: 2});
                }
            }
        });
    },
    // 取消赞
    cancelZan: function (elem, type, id) {
        jQuery.ajax({
            type: 'POST',
            url:  '/interact/cancel-zan',
            data: {
                id:   id,
                type: type
            },
            sync: true,
            success: function (r, status) {
                if (r.code == 0) {
                    let number = $(elem).find('.number').html()
                    $(elem).find('.number').html(parseInt(number) - 1)
                    $(elem).find('.icon').removeClass('icon-zan').addClass('icon-zan1')
                } else {
                    layer.alert(r.message, {icon: 2});
                }
            }
        });
    },
    // 检查是执行踩还是取消踩
    checkCai: function (elem, id) {
        var type = $(elem).attr("data-type")
        if ($(elem).find('.icon').hasClass('icon-cai1')) {
            this.cancelCai(elem, type, id)
        } else {
            this.cai(elem, type, id)
        }
    },
    // 踩
    cai: function (elem, type, id) {
        jQuery.ajax({
            type: 'POST',
            url : '/interact/cai',
            data: {
                id:   id,
                type: type
            },
            sync: true,
            success: function (r, status) {
                if (r.code == 0) {
                    let number = $(elem).find('.number').html()
                    $(elem).find('.number').html(parseInt(number)+1)
                    $(elem).find('.icon').removeClass('icon-cai2').addClass('icon-cai1')
                } else {
                    layer.alert(r.message, {icon: 2});
                }
            }
        });
    },
    // 取消踩
    cancelCai: function (elem, type, id) {
        jQuery.ajax({
            type: 'POST',
            url:  '/interact/cancel-cai',
            data: {
                id:   id,
                type: type
            },
            sync: true,
            success: function (r, status) {
                if (r.code == 0) {
                    let number = $(elem).find('.number').html()
                    $(elem).find('.number').html(parseInt(number) - 1)
                    $(elem).find('.icon').removeClass('icon-cai1').addClass('icon-cai2')
                } else {
                    layer.alert(r.message, {icon: 2});
                }
            }
        });
    }
}

// 内容模块
gf.content = {
    // 删除内容
    delete: function (id, url) {
        url = url || "/"
        layer.confirm('您确定要删除该内容吗？', {icon: 0, title:'删除内容', offset: '40%'}, function(index){
            jQuery.ajax({
                type: 'POST',
                url: '/content/do-delete',
                data: {
                    id: id,
                },
                success: function (r) {
                    if (r.code === 0) {
                        layer.alert(r.message, {icon: 1, offset: '40%'}, function(){
                            window.location.href = url;
                        });
                    } else {
                        layer.alert(r.message, {icon: 2, offset: '40%'});
                    }
                },
            });
        });
    }
}


