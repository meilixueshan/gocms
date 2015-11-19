<!DOCTYPE html>
<html ng-app="cms">

<head>
    <title>单页面列表</title>
    {{template "admin/resources.tpl" .}}
</head>

<body ng-controller="pageListController">
    <form id="form1" name="form1" method="post" action="" class="xw_container">
        <table border="0" cellspacing="1" cellpadding="0" class="xw_list_table">
            <tr>
                <th style="width:160px;">ID</th>
                <th style="width:240px;">标题</th>
                <th>SEO标题</th>
                <th style="width:60px;">管理</th>
            </tr>
			<tr ng-repeat="item in data">
				<td>{$item.id$}</td>
				<td>{$item.title$}</td>
				<td>{$item.seotitle$}</td>
				<td><a href="/admin/page/edit/{$item.id$}">修改</a>
					<a href="javascript:void(0);" ng-click="del(item.id)">删除</a>
				</td>
			</tr>
        </table>
    </form>
</body>

</html>