<!DOCTYPE html>
<html ng-app="cms">

<head>
    <title>单页面列表</title>
    {{template "admin/resources.tpl" .}}
</head>

<body ng-controller="pageController">
    <form id="form1" name="form1" method="post" action="" class="xw_container">
        <table border="0" cellspacing="1" cellpadding="0" class="xw_edit_table">
            <tr>
                <td class="label">页面名称：</td>
                <td class="input">
                    <input type="text" id="txtId" ng-model="data.Id" class="win" value="{{.id}}" />
                </td>
            </tr>
            <tr>
                <td class="label">页面标题：</td>
                <td class="input">
                    <input type="text" id="txtTitle" ng-model="data.Title" class="win" />
                </td>
            </tr>
            <tr>
                <td class="label">页面内容：</td>
                <td class="input">
                    <textarea id="txtContent" name="txtContent" class="wtextarea" ng-model="data.Content"></textarea>
                </td>
            </tr>
            <tr>
                <td class="label">SEO标题：</td>
                <td class="input">
                    <input type="text" id="txtSEOTitle" ng-model="data.Seotitle" class="win" />
                </td>
            </tr>
            <tr>
                <td class="label">SEO关键词：</td>
                <td class="input">
                    <input type="text" id="txtSEOKeywords" ng-model="data.Seokeywords" class="win" />
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
                    <input type="button" id="btnSave" name="btnSave" value="保存" class="xw_btn" ng-click="save(data)" />
                </td>
            </tr>
        </table>
    </form>
</body>

</html>