<!DOCTYPE html>
<html ng-app="cms">

<head>
    <title>基本信息</title>
    {{template "admin/resources.tpl" .}}
</head>

<body ng-controller="siteController">
    <form id="form1" name="form1" method="post" action="" class="xw_container">
        <table border="0" cellspacing="1" cellpadding="0" class="xw_edit_table">
            <tr>
                <td class="label">域名：</td>
                <td class="input">
                    <input type="text" id="txtDomainName" class="win" value="" maxlength="255" ng-model="siteData.DomainName" />
                </td>
            </tr>
            <tr>
                <td class="label">网站名称：</td>
                <td class="input">
                    <input type="text" id="txtSiteName" class="win" value="" maxlength="255" ng-model="siteData.SiteName" />
                </td>
            </tr>
            <tr>
                <td class="label">公司名称：</td>
                <td class="input">
                    <input type="text" id="txtCompanyName" class="win" value="" maxlength="255" ng-model="siteData.CompanyName" />
                </td>
            </tr>
            <tr>
                <td class="label">地址：</td>
                <td class="input">
                    <input type="text" id="txtAddress" class="win" value="" maxlength="255" ng-model="siteData.Address" />
                </td>
            </tr>
            <tr>
                <td class="label">电话：</td>
                <td class="input">
                    <input type="text" id="txtTel" class="win" value="" maxlength="255" ng-model="siteData.Tel" />
                </td>
            </tr>
            <tr>
                <td class="label">传真：</td>
                <td class="input">
                    <input type="text" id="txtFax" class="win" value="" maxlength="255" ng-model="siteData.Fax" />
                </td>
            </tr>
            <tr>
                <td class="label">联系人：</td>
                <td class="input">
                    <input type="text" id="txtLinkman" class="win" value="" maxlength="255" ng-model="siteData.Linkman" />
                </td>
            </tr>
            <tr>
                <td class="label"></td>
                <td class="input">
                    <input type="button" id="btnSave" name="btnSave" value="保存" class="xw_btn" ng-click="save(siteData)" />
                </td>
            </tr>
        </table>
    </form>
</body>

</html>