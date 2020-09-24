$(function () {
    $('#send-comment').on('click', function () {
        var form = $('.comment-content-form');
        if (form.css('display') === 'none') {
            form.css('display', 'block');
            $("html,body").animate({scrollTop: $(".comment-content-form").offset().top}, 1000)
        } else {
            form.css('display', 'none')
        }
    });
    // $('#submit-comment').on('click', function () {
    //     var name = $('#name').val(), email = $('#email').val(), site = $('#site').val(), content = $('#content').val();
    //     if (!name || !content) {
    //         alert('*号为必填项哦');
    //         return
    //     }
    //     $('#submit-comment').attr('disabled', true);
    //     $.ajax({
    //         url: "/articles/" + window.articleId + "/comments",
    //         type: "POST",
    //         data: {name: name, email: email, site: site, content: content, avatar: $('.comment-user-avatar').val()},
    //         dataType: 'json',
    //         success: function (res) {
    //             if (res.status) {
    //                 getArticleComments(5, 0, true)
    //             }
    //             alert(res.message);
    //             $('#submit-comment').attr('disabled', false)
    //         },
    //         error: function (res) {
    //             alert('留言失败')
    //         }
    //     })
    // });
    //
    // function getArticleComments(limit = 5, offset = 0, fresh = false) {
    //     $('.load-more-comments').html('loading...');
    //     setTimeout(function () {
    //         $.ajax({
    //             url: "/articles/" + window.articleId + "/comments",
    //             type: "GET",
    //             data: {limit: limit, offset: offset},
    //             setTimeout(handler, timeout, ...arguments) {
    //             },
    //             dataType: 'json',
    //             success: function (res) {
    //                 var d = '加载更多';
    //                 if (res == '') {
    //                     d = '无更多数据'
    //                 } else {
    //                     fresh ? $('.comment-list').html(res) : $('.comment-list').append(res)
    //                 }
    //                 $('.load-more-comments').html(d);
    //                 $('.footer').show()
    //             },
    //             error: function (res) {
    //                 alert('留言失败')
    //             }
    //         })
    //     }, 300)
    // }
    //
    // var limit = 5, offset = 0;
    // getArticleComments(limit, offset, false);
    // $('.load-more-comments').on('click', function () {
    //     offset = $('.comment-list .comment-item').length;
    //     if (offset % limit == 0) {
    //         getArticleComments(limit, offset, false)
    //     } else {
    //         $('.load-more-comments').html('无更多数据')
    //     }
    // });
    // $('#email').on('input', function () {
    //     $.ajax({
    //         url: "/avatar", type: "GET", data: {email: $(this).val()}, dataType: 'json', success: function (res) {
    //             $('.comment-avatar').attr('src', res.url);
    //             $('.comment-user-avatar').val(res.url)
    //         }, error: function (res) {
    //         }
    //     })
    // })
});