<!DOCTYPE html>
<html ng-app="cms">

<head>
    <title>文章分类列表</title>
    {{template "admin/resources.tpl" .}}
</head>

<body ng-controller="articleCategoryListController">
    <form id="form1" name="form1" method="post" action="" class="xw_container">
        <table border="0" cellspacing="1" cellpadding="0" class="xw_list_table">
            <tr>
                <th style="width:30px;">编号</th>
                <th>分类名称</th>
                <th style="width:140px;">级联编码</th>
                <th style="width:60px;">排序编号</th>
                <th style="width:60px;">管理</th>
            </tr>
            {{range $key, $val := .categories}}
			<tr>
				<td>{{$val.Id}}</td>
				<td style="text-align:left">{{CategoryFormat $val.Categoryname $val.Innercode}}</td>
				<td style="text-align:left">{{$val.Innercode}}</td>
				<td class="tdSortNumber" data-id="{{$val.Id}}">{{$val.Sortnum}}</td>
				<td><a href="/admin/articlecategory/edit/{{$val.Id}}">修改</a>
					<a href="javascript:void(0)" ng-click="del({{$val.Id}})">删除</a>
				</td>
			</tr>
            {{end}}
        </table>
    </form>
</body>

</html>