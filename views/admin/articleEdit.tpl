<!DOCTYPE html>
<html ng-app="cms">

<head>
    <title>文章编辑</title>
    {{template "admin/resources.tpl" .}}
</head>

<body ng-controller="articleController">
    <form id="form1" name="form1" method="post" action="" class="xw_container">
        <table border="0" cellspacing="1" cellpadding="0" class="xw_edit_table">
            <tr>
                <td class="label">文章标题：</td>
                <td class="input">
                    <input type="text" id="txtTitle" class="win" ng-model="data.Title" />
                </td>
            </tr>
            <tr>
                <td class="label">文章摘要：</td>
                <td class="input">
                    <textarea id="txtDescription" class="wtextarea" ng-model="data.Description"></textarea>
                </td>
            </tr>
            <tr>
                <td class="label">所属分类：</td>
                <td class="input">
					<select id="drpCategory" ng-model="data.Categoryid">
                        <option value="">───请选择分类───</option>
						{{range $key, $val := .categories}}
						<option value="{{$val.Id}}">{{CategoryFormat $val.Categoryname $val.Innercode}}</option>
						{{end}}
						<option value="8">新闻中心</option>
                    </select>
                </td>
            </tr>
            <tr>
                <td class="label">封面图片：</td>
                <td class="input">
                    <input type="text" id="txtPicFileName" class="win" ng-model="data.Picfilename" />
                </td>
            </tr>
            <tr>
                <td class="label">文章作者：</td>
                <td class="input">
                    <input type="text" id="txtAuthor" class="win" ng-model="data.Author" />
                </td>
            </tr>
            <tr>
                <td class="label">文章来源：</td>
                <td class="input">
                    <input type="text" id="txtComeFrom" class="win" ng-model="data.Comefrom" />
                </td>
            </tr>
            <tr>
                <td class="label">发布时间：</td>
                <td class="input">
                    <input type="text" id="txtPublishTime" class="win" ng-model="data.Publishtime" />
                </td>
            </tr>
            <tr>
                <td class="label">阅读次数：</td>
                <td class="input">
                    <input type="text" id="txtReadNum" class="win" ng-model="data.Readnum" />
                </td>
            </tr>
            <tr>
                <td class="label">文章内容：</td>
                <td class="input">
                    <textarea id="txtContent" name="txtContent" style="width:100%" ng-model="data.Content"></textarea>
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
                    <input type="button" id="btnSave" value="保存" class="xw_btn" ng-click="save(data)" />
                    <input type="hidden" id="hideId" ng-model="data.Id" value="{{.id}}" />
                </td>
            </tr>
        </table>
    </form>
</body>

</html>