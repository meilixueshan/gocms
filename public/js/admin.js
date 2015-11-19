/*********************************************************************************
 ***分页信息***********************************************************************
 *********************************************************************************/
function division(value1, value2) {
    var n1 = Math.round(value1); //四舍五入  
    var n2 = Math.round(value2); //四舍五入  

    var result = n1 / n2; //除  

    if (result >= 0) {
        result = Math.floor(result); //返回小于等于原rslt的最大整数。  
    } else {
        result = Math.ceil(result); //返回大于等于原rslt的最小整数。  
    }

    return result;
}
function simplePager(pageSize, pageNum, recordCount) {
	var maxPage = 1;
    if (pageSize < 1) {
        pageSize = 1;
    }

    if (recordCount > 0) {
        if (recordCount % pageSize == 0) {
            maxPage = division(recordCount, pageSize);
        } else {
            maxPage = division(recordCount, pageSize) + 1;
        }
    }

    if (pageNum < 1) {
        pageNum = 1;
    }

    if (pageNum > maxPage) {
        pageNum = maxPage;
    }

    var output = "";
    if (pageNum <= 1) {
        output += '<span>首页</span>';
        output += '<span>上页</span>';
    } else {
        output += '<span><a ng-click="search({0})">首页</a></span>'._format(1);
        output += '<span><a ng-click="search({0})">上页</a></span>'._format(pageNum - 1);
    }

    if (pageNum >= maxPage) {
        output += '<span>下页</span>';
        output += '<span>尾页</span>';
    } else {
        output += '<span><a ng-click="search({0})">下页</a></span>'._format(pageNum + 1);
        output += '<span><a ng-click="search({0})">尾页</a></span>'._format(maxPage);
    }

    return output;
}

/*********************************************************************************
 ***控制器***********************************************************************
 *********************************************************************************/
angular.module('cms', [])
.config(function($interpolateProvider) {
    $interpolateProvider.startSymbol('{$');
    $interpolateProvider.endSymbol('$}');
})
.controller("siteController", function($scope, $http) {
	$http.get('/admin/site/getdata').success(function(data) {
		$scope.siteData = data;
	}).error(function(data, status, headers, config) {
        _mapp.alert(data);
    });
	
	$scope.save = function(siteData) {
		console.log(siteData);
		_mapp.loading();
		$http.post('/admin/site/save', siteData).success(function(data) {
			if(data.Success) {
				_mapp.msg("保存成功！");
			} else {
				_mapp.alert(data.error);
			}
			_mapp.loadOver();
		});
	};
})
.controller("pwdController", function($scope, $http) {
	$scope.data = {
		"oldpwd": "",
		"newpwd": "",
		"newpwd2": ""
	};
	$scope.save = function(uiData) {
		if(uiData.oldpwd === "") {
			_mapp.alert("请输入原密码！");
			return;
		}
		if(uiData.newpwd === "") {
			_mapp.alert("请输入新密码！");
			return;
		}
		if(uiData.newpwd != uiData.newpwd2) {
			_mapp.alert("两次输入的新密码不同！");
			return;
		}
		_mapp.loading();
		$http.post('/admin/password/save', uiData).success(function(data) {
			if(data.Success) {
				uiData.oldpwd = "";
				uiData.newpwd = "";
				uiData.newpwd2 = "";
				
				_mapp.msg("保存成功！");
			} else {
				_mapp.alert(data.Error);
			}
			_mapp.loadOver();
		}).error(function(data, status, headers, config) {
        	_mapp.alert(data);
    	});
	};
})
.controller("pageListController", function($scope, $http) {
	$http.get('/admin/page/getdatalist').success(function(data){
		$scope.data = data;
	});
	
	$scope.del = function(id) {
		_mapp.confirm("确定要删除此信息吗？", function() {
			$http.get('/admin/page/delete/' + id).success(function(result) {
				if(result.Success) {
					_mapp.msg("删除成功！");
					window.location.reload(0);
				} else {
					_mapp.msg(result.Error);
				}
			}).error(function(result) {
				_mapp.alert(result);
			});	
		});
	};
})
.controller("pageController", function($scope, $http) {
	var pageId = $("#txtId").val();
	if(pageId !== "") {
		var url = '/admin/page/getdata/' + pageId;
		$http.get(url).success(function(data) {
			$scope.data = data;
		}).error(function(data, status, headers, config) {
	        _mapp.alert(data);
	    });		
	}
	
	$scope.save = function(uiData) {
		_mapp.loading();
		$http.post("/admin/page/save", uiData).success(function(result) {
			if(result.Success) {
				window.location = "/admin/page";
			} else {
				_mapp.loadOver();
				_mapp.alert(result.Error);
			}
		}).error(function(result) {
			_mapp.loadOver();
			_mapp.alert(result);
		});
	};
})
.controller("articleListController", function($scope, $http, $sce, $compile) {
	$scope.search = function(page) {
		if(page) {
			currentPage = page.toString();
		} else {
			currentPage = "1";
		}

		var jsonObj = {
			pageSize : "10",
			pageNo : currentPage,
			categoryId : $("#drpCategory").val(),
			keywords : $("#txtKewords").val()
		};
		$http.post('/admin/article/getdatalist', jsonObj).success(function(result){
			$scope.data = result.data;
			$scope.totalCount = result.totalCount;
			
			//var html = $sce.trustAsHtml(simplePager(parseInt(jsonObj.pageSize), parseInt(jsonObj.pageNo), result.totalCount, ""));
			var html = simplePager(parseInt(jsonObj.pageSize), parseInt(jsonObj.pageNo), result.totalCount, "");
			var template = angular.element(html);
			var element = $compile(template)($scope);
			angular.element("#pager").html(element);
		});
	}
	
	$scope.search();
	
	$scope.del = function(id) {
		_mapp.confirm("确定要删除此信息吗？", function() {
			$http.get('/admin/article/delete/' + id).success(function(result) {
				if(result.Success) {
					_mapp.msg("删除成功！");
					window.location.reload(0);
				} else {
					_mapp.msg(result.Error);
				}
			}).error(function(result) {
				_mapp.alert(result);
			});	
		});
	};
})
.controller("articleController", function($scope, $http) {
	var articleId = $("#hideId").val();
	if(articleId !== "") {
		var url = '/admin/article/getdata/' + articleId;
		$http.get(url).success(function(data) {
			$scope.data = data;
		}).error(function(data, status, headers, config) {
	        _mapp.alert(data);
	    });		
	} else {
		var currentTime = new Date()._format("yyyy-MM-dd hh:mm:ss");
		$scope.data = {"Publishtime": currentTime};
	}
	
	$scope.save = function(uiData) {
		uiData.Id = parseInt(uiData.Id);
		uiData.Categoryid = parseInt(uiData.Categoryid);
		uiData.Readnum = parseInt(uiData.Readnum);
		if(uiData.Categoryid <= 0) {
			_mapp.msg("请选择所属分类！");
			return;
		}
		
		_mapp.loading();
		$http.post("/admin/article/save", uiData).success(function(result) {
			if(result.Success) {
				window.location = "/admin/article";
			} else {
				_mapp.loadOver();
				_mapp.alert(result.Error);
			}
		}).error(function(result) {
			_mapp.loadOver();
			_mapp.alert(result);
		});
	};
})
.controller("articleCategoryListController", function($scope, $http) {
	$scope.del = function(id) {
		_mapp.confirm("确定要删除此信息吗？", function() {
			$http.get('/admin/articlecategory/delete/' + id).success(function(result) {
				if(result.Success) {
					_mapp.msg("删除成功！");
					window.location.reload(0);
				} else {
					_mapp.msg(result.Error);
				}
			}).error(function(result) {
				_mapp.alert(result);
			});	
		});
	};
})
.controller("articleCategoryContoller", function($scope, $http) {
	var categoryId = $("#hideId").val();
	if(categoryId !== "") {
		var url = '/admin/articlecategory/getdata/' + categoryId;
		$http.get(url).success(function(data) {
			$scope.data = data;
		}).error(function(data, status, headers, config) {
	        _mapp.alert(data);
	    });		
	} else {
		$scope.data = { "Parentid": 0, "Sortnum":99 };
	}
	
	$scope.save = function(uiData) {
		uiData.Id = parseInt(uiData.Id);
		uiData.Parentid = parseInt(uiData.Parentid);
		uiData.Sortnum = parseInt(uiData.Sortnum);

		_mapp.loading();
		$http.post("/admin/articlecategory/save", uiData).success(function(result) {
			if(result.Success) {
				window.location = "/admin/articlecategory";
			} else {
				_mapp.loadOver();
				_mapp.alert(result.Error);
			}
		}).error(function(result) {
			_mapp.loadOver();
			_mapp.alert(result);
		});
	};
});