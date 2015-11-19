<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>GoCMS</title>
	<link rel="stylesheet" href="/public/lib/pure/pure-min.css" />
	<link rel="stylesheet" href="/public/css/main.css" />
</head>
<body>
{{template "front/header.tpl" .}}
<div id="layout" class="pure-g">
	<div class="pure-g">
	    <div class="pure-u-5-5">
			<div class="main">
	            <!-- A wrapper for all the blog posts -->
				{{range $key, $val := .articles}}
	            <div class="posts">
	                <!-- A single blog post -->
	                <section class="post">
	                    <header class="post-header">
	                        <h2 class="post-title"><a href="/article/show/{{$val.id}}">{{$val.title}}</a></h2>
	
	                        <p class="post-meta">
	                            <a href="/article/list/{{$val.categoryid}}" class="post-author">[{{$val.categoryname}}]</a> Created {{$val.publishtime}} By {{$val.author}} Read {{$val.readnum}}
	                        </p>
	                    </header>
	
	                    <div class="post-description">
	                        <p>{{$val.description}}</p>
	                    </div>

						<div class="post-footer">
						</div>
	                </section>
	            </div>
				{{end}}
	
	            {{template "front/footer.tpl" .}}
	        </div>
		</div>
	</div>
</div>
</body>
</html>