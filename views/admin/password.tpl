<!DOCTYPE html>
<html ng-app="cms">

<head>
    <title>修改密码</title>
    {{template "admin/resources.tpl" .}}
</head>

<body ng-controller="pwdController">
    <form id="form1" name="form1" method="post" action="" class="xw_container">
        <table border="0" cellspacing="1" cellpadding="0" class="xw_edit_table">
            <tr>
                <td class="label">原密码：</td>
                <td class="input">
                    <input type="password" id="oldPassword" name="oldPassword" class="win" value="" ng-model="data.oldpwd" />
                </td>
            </tr>
            <tr>
                <td class="label">新密码：</td>
                <td class="input">
                    <input type="password" id="newPassword" name="newPassword" class="win" value="" ng-model="data.newpwd" />
                </td>
            </tr>
            <tr>
                <td class="label">确认新密码：</td>
                <td class="input">
                    <input type="password" id="newPassword2" name="newPassword2" class="win" value="" ng-model="data.newpwd2" />
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