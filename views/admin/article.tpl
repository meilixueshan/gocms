<!DOCTYPE html>
<html ng-app="cms">

<head>
    <title>文章列表</title>
    {{template "admin/resources.tpl" .}}
</head>

<body ng-controller="articleListController as myctl">
    <form id="form1" name="form1" method="post" action="" class="xw_container">
		<table border="0" cellspacing="0" cellpadding="0" align="center" class="xw_table_search">
            <tr>
                <td>
					分类：<select id="drpCategory">
					<option value="">───所有分类───</option>
					{{range $key, $val := .categories}}
					<option value="{{$val.Id}}">{{CategoryFormat $val.Categoryname $val.Innercode}}</option>
					{{end}}
					</select>　　
					关键词：
                    <input type="text" id="txtKewords" value="" class="txt" />
					<input type="button" id="btnSearch" value="搜索" class="xw_btn" ng-click="search()" />
                </td>
            </tr>
        </table>
        <table border="0" cellspacing="1" cellpadding="0" class="xw_list_table">
            <tr>
                <th style="width:40px;">编号</th>
                <th style="width:140px;">分类名称</th>
                <th>标题</th>
                <th style="width:120px;">发布时间</th>
                <th style="width:120px;">创建时间</th>
                <th style="width:60px;">管理</th>
            </tr>
			<tr ng-repeat="item in data">
				<td>{$item.id$}</td>
				<td>{$item.categoryname$}</td>
				<td>{$item.title$}</td>
				<td>{$item.publishtime$}</td>
				<td>{$item.createtime$}</td>
				<td><a href="/admin/article/edit/{$item.id$}">修改</a>
					<a href="javascript:void(0);" ng-click="del(item.id)">删除</a>
				</td>
			</tr>
        </table>
        <table border="0" cellspacing="0" cellpadding="0" align="center" class="xw_table_pager">
            <tr>
                <td><div id="pager"></div>
				</td>
            </tr>
        </table>
    </form>
</body>

</html>