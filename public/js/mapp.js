var _mapp = {
    alert: function (content, callbackFn) {
        layer.alert(content, function (index) {
            if (callbackFn) {
                callbackFn();
            }
            layer.close(index);
        });
    },
    confirm: function (msg, callbackFn) {
        layer.confirm(msg, {
            icon: 3
        }, function (index) {
            layer.close(index);
            callbackFn();
        });
    },
	msg: function(msg) {
		layer.msg(msg);
	},
	loading: function() {
		var index = layer.load(1, {
		    shade: [0.1,'#000'] //0.1透明度的白色背景
		});
		return index;
	},
	loadOver: function() {
		layer.closeAll('loading');
	},
    openWin: function (title, url, width, height) {
        layer.open({
            type: 2,
            title: title,
            shade: 0.5,
            skin: 'layui-layer-rim', //加上边框
            area: [width, height],
            content: [url, 'no']
        });
    }
};

/*********************************************************************************
 ***原形扩展***********************************************************************
 *********************************************************************************/
String.prototype._format = function () {
    if (arguments.length == 0) return this;
    for (var s = this, i = 0; i < arguments.length; i++)
        s = s.replace(new RegExp("\\{" + i + "\\}", "g"), arguments[i]);
    return s;
};

Date.prototype._format = function (fmt) {
    var o = {
        "M+": this.getMonth() + 1, //月份 
        "d+": this.getDate(), //日 
        "h+": this.getHours(), //小时 
        "m+": this.getMinutes(), //分 
        "s+": this.getSeconds(), //秒 
        "q+": Math.floor((this.getMonth() + 3) / 3), //季度 
        "S": this.getMilliseconds() //毫秒 
    };
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
    if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;
};

String.prototype._replace = function (s1, s2) {
    return this.replace(new RegExp(s1, "gm"), s2);
};

String.prototype._startWith = function (str) {
    if (str == null || str == "" || this.length == 0 || str.length > this.length)
        return false;

    return (this.substr(0, str.length) == str);
};

String.prototype._endWith = function (str) {
    if (str == null || str == "" || this.length == 0 || str.length > this.length)
        return false;

    return (this.substring(this.length - str.length) == str);
};

/*********************************************************************************
 ***Ajax简化***********************************************************************
 *********************************************************************************/
$(function () {
    /**
     * ajax封装
     * url 发送请求的地址
     * data 发送到服务器的数据，数组存储，如：{"date": new Date().getTime(), "state": 1}
     * async 默认值: true。默认设置下，所有请求均为异步请求。如果需要发送同步请求，请将此选项设置为 false。
     *       注意，同步请求将锁住浏览器，用户其它操作必须等待请求完成才可以执行。
     * type 请求方式("POST" 或 "GET")， 默认为 "GET"
     * dataType 预期服务器返回的数据类型，常用的如：xml、html、json、text
     * successfn 成功回调函数
     * errorfn 失败回调函数
     */
    jQuery.ajaxHttp = function (url, data, async, type, dataType, successfn, errorfn) {
        async = (async == null || async == "" || typeof (async) == "undefined") ? "true" : async;
        type = (type == null || type == "" || typeof (type) == "undefined") ? "post" : type;
        dataType = (dataType == null || dataType == "" || typeof (dataType) == "undefined") ? "json" : dataType;
        data = (data == null || data == "" || typeof (data) == "undefined") ? {
            "date": new Date().getTime()
        } : data;
        $.ajax({
            type: type,
            async: async,
            data: data,
            url: url,
            dataType: dataType,
            success: function (d) {
                successfn(d);
            },
            error: function (a, e) {
                errorfn(e);
            }
        });
    };

    /**
     * ajax封装
     * url 发送请求的地址
     * data 发送到服务器的数据，数组存储，如：{"date": new Date().getTime(), "state": 1}
     * type POST/GET/PUT/DELETE
     * successfn 成功回调函数
     * errorfn 失败回调函数
     */
    jQuery.ajaxDo = function (url, data, type, successfn, errorfn) {
        data = (data == null || data == "" || typeof (data) == "undefined") ? {
            "date": new Date().getTime()
        } : data;
        $.ajax({
            type: type,
            data: data,
            url: url,
            dataType: "text",
            success: function (d) {
                successfn(d);
            },
            error: function (a, e) {
                errorfn(e);
            }
        });
    };

    $(".xw_list_table tr td").bind("mouseover", function () {
        $(this).parent().children("td").css({
            background: "#e5e5e5"
        });
    });
    $(".xw_list_table tr td").bind("mouseout", function () {
        $(this).parent().children("td").css({
            background: "#ffffff"
        });
    });
});