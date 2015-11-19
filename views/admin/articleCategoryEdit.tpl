<!DOCTYPE html>
<html ng-app="cms">

<head>
    <title>文章分类编辑</title>
    {{template "admin/resources.tpl" .}}
</head>

<body ng-controller="articleCategoryContoller">
    <form id="form1" name="form1" method="post" action="" class="xw_container">
        <table border="0" cellspacing="1" cellpadding="0" class="xw_edit_table">
            <tr>
                <td class="label">分类名称：</td>
                <td class="input">
                    <input type="text" id="txtCategoryName" class="win" value="" ng-model="data.Categoryname" />
                </td>
            </tr>
            <tr>
                <td class="label">父级分类：</td>
                <td class="input">
                    <select id="drpParent" ng-model="data.Parentid">
                        <option value="0">---根分类---</option>
						{{range $key, $val := .categories}}
						<option value="{{$val.Id}}">{{CategoryFormat $val.Categoryname $val.Innercode}}</option>
						{{end}}
						</select>　
                    </select>
                </td>
            </tr>
            <tr>
                <td class="label">排序编号：</td>
                <td class="input">
                    <input type="text" id="txtSortNum" class="win" value="" ng-model="data.Sortnum" />
                </td>
            </tr>
            <tr>
                <td class="label">SEO标题：</td>
                <td class="input">
                    <input type="text" id="txtSEOTitle" class="win" value="" ng-model="data.Seotitle" />
                </td>
            </tr>
            <tr>
                <td class="label">SEO关键词：</td>
                <td class="input">
                    <input type="text" id="txtSEOKeywords" class="win" value="" ng-model="data.Seokeywords" />
                </td>
            </tr>
            <tr>
                <td class="label">SEO描述：</td>
                <td class="input">
                    <textarea id="txtSEODescription" class="wtextarea" ng-model="data.Seodescription"></textarea>
                </td>
            </tr>
            <tr>
                <td class="label"></td>
                <td class="input">
                    <input type="button" id="btnSave" value="保存" class="xw_btn" ng-click="save(data)" />
                    <input type="hidden" id="hideId" ng-model="data.Id" value="{{.id}}" />
                </td>
            </tr>
        </table>
    </form>
</body>

</html>