<!DOCTYPE html>
<html>
<head>
    <title>系统后台登录</title>
	{{template "admin/resources.tpl" .}}
	<link rel="stylesheet" type="text/css" href="/public/css/admin_m.css" />
	<script type="text/javascript">
	function setTab(name,cursel,n){
		for(i=1;i<=n;i++){
			var menu=document.getElementById(name+i);
			var con=document.getElementById("con_"+name+"_"+i);
			try {
			menu.className=i==cursel?"navon":"";
			con.style.display=i==cursel?"block":"none";
			}catch(e){}
		}
		return false;
	}
	</script> 
	<script language="javascript">
		document.onselectstart=mylock1;
		document.oncontextmenu=mylock1;
		function mylock1(){
		  event.returnValue=false;
		}
	</script>
</head>
<body style="height:100%; overflow:hidden;">
	<table border="0" cellpadding="0" cellspacing="0" height="100%" width="100%">
        <tbody>
            <tr>
                <td colspan="2" height="77" valign="top">
                    <div id="header">
                        <div class="logo fl">
                            <div class="png">
                                <a href="../"><img src="/public/images/xw_loginlogo.png" class="ie6" height="25" width="143"></a>
                            </div>
                            <div class="lun">
                                <span style="color:#FA891B">{{.appname}} V1.0.0</span>(20151101)
                            </div>
                        </div>
                        <ul class="nav">
                            <li id="nav1" onclick="return setTab('nav',1,5)" class="navon"><em><a href="#">常用功能</a></em></li>
                            <li id="nav2" onclick="return setTab('nav',2,5)" style="display:none"><em><a href="#">定制模块</a></em></li>
                            <li id="nav3" onclick="return setTab('nav',3,5)" style="display:none"><em><a href="#">人才招聘</a></em></li>
                            <li id="nav4" onclick="return setTab('nav',4,5)" style="display:none"><em><a href="#">在线留言</a></em></li>
                            <li id="nav5" onclick="return setTab('nav',5,5)" style="display:none"><em><a href="#">文件下载</a></em></li>
                        </ul>
                        <div class="wei2">
                            <div class="wei fl">
                                <span style=" float:left">当前用户：{{.user.Username}}&nbsp;|&nbsp;
								<a href="/admin/logout" class="loginOut" title="退出后台">退出后台</a>
							</span> &nbsp;|&nbsp;
                                <a target="main" href="site">基本信息</a> &nbsp;|&nbsp;
                                <a target="main" href="password">修改密码</a> &nbsp;|&nbsp;
                                <a href="../" style="cursor: pointer;" class="siteIndex" target="_blank">网站首页</a>
                            </div>
                        </div>
                    </div>
                </td>
            </tr>
            <tr>
                <td id="main-fl" style="height:94%; overflow:hidden; " valign="top">
                    <div id="left" style="height:100%; overflow:auto;">
                        <div id="con_nav_1">
                            <h1>文章模块</h1>
                            <div class="cc"></div>
                            <ul>
                                <li><a target="main" href="articlecategory/">文章分类</a></li>
                                <li><a target="main" href="articlecategory/edit">新增分类</a></li>
                                <li><a target="main" href="article/">文章列表</a></li>
                                <li><a target="main" href="article/edit">新增文章</a></li>
                            </ul>
                            <h1>单页面模块</h1>
                            <div class="cc"></div>
                            <ul>
                                <li><a target="main" href="/admin/page">管理单页面</a></li>
                                <li><a target="main" href="/admin/page/edit">新增单页面</a></li>
                            </ul>
                        </div>
                        <div id="con_nav_2" style="display:none;">
                            <h1>定制模块</h1>
                            <div class="cc"></div>
                            <ul>
                                <li><a target="main" href="/admin/imgtxtposition">位置列表</a></li>
                                <li><a target="main" href="/admin/imgtxtposition/edit">新增位置</a></li>
                            </ul>
                            <h1>关键词</h1>
                            <div class="cc"></div>
                            <ul>
                                <li><a target="main" href="/admin/seo">关键词列表</a></li>
                                <li><a target="main" href="/admin/seo/edit">新增关键词</a></li>
                            </ul>
                        </div>
                        <div id="con_nav_3" style="display:none;">
                            <h1>人才招聘</h1>
                            <div class="cc"></div>
                            <ul>
                                <li><a target="main" href="#">招聘职位</a></li>
                                <li><a target="main" href="#">新增职位</a></li>
                            </ul>
                        </div>
                        <div id="con_nav_4" style="display:none;">
                            <h1>在线留言</h1>
                            <div class="cc"></div>
                            <ul>
                                <li><a target="main" href="#">未回复留言</a></li>
                                <li><a target="main" href="#">已回复留言</a></li>
                            </ul>
                        </div>
                        <div id="con_nav_5" style="display:none;">
                            <h1>文件下载</h1>
                            <div class="cc"></div>
                            <ul>
                                <li><a target="main" href="#">文件分类</a></li>
                                <li><a target="main" href="#">新增分类</a></li>
                                <li><a target="main" href="#">文件列表</a></li>
                                <li><a target="main" href="#">新增文件</a></li>
                            </ul>
                        </div>
                    </div>
                </td>
                <td id="mainright" style="height:94%; " valign="top">
                    <iframe id="frmMain" name="main" scrolling="yes" style="overflow: visible;" src="" frameborder="0" height="100%" width="100%"></iframe>
                </td>
            </tr>
        </tbody>
    </table>
</body>
</html>