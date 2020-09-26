package handler

import "text/template"

const (
	homeTpl = `
	<!DOCTYPE html>
	<html>
	
	<head>
			<script src="//code.jquery.com/jquery-1.11.3.min.js"></script>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
		<title>文章列表</title>
		<style>
	
			.main {
				margin: 0 auto;
				width: 400px;
				border: 1px solid;
				border-color: #eeeeee;
				border-radius: 5px;
				margin-top: 100px;
			}
		</style>
	
	</head>
	<body>
			<div data-v-0c10c649="" class="course-tab-view" >
					{{range $idx, $ele := .}}
					<div data-v-0c10c649="" class="table-item-wrap">
						<div data-v-0c10c649="" class="table-item">
						  <div data-v-0c10c649="" class="table-item-text">
							<div data-v-0c10c649="" class="table-item-text-subtitle" id="{{$ele.Class}}" fpath="{{$ele.Path}}"  >{{$ele.Name}}</div></div>
						  <div data-v-0c10c649="" class="table-item-right">
							<i data-v-0c10c649="" class="iconfont lock"></i></div>
						</div>
					</div>
					{{end}}
				  </div>
	</body>
	<script>
		$(function () {
		// init...
		regEvent()
	});
	</script>
	<script>
	function regEvent(){
		console.log("in reg")
		var list = $('.table-item-text-subtitle')
		list.click(handleClick)
	}
	</script>
	
	<script>
			function handleClick(e){
				var obj = $(this)
				, id = obj.attr("id")
				, fpath = obj.attr("fpath")
				, name = obj.text()
				, url = window.location.href + '?id=' + id;
				// console.log(name)
				url = "/content/" + name + '?id=' + id + '&?fpath=' + fpath;
				window.location.href = url
			}
			</script>
	</html>
	`
)

var tpls = map[string]*template.Template{}

func init() {
	m := map[string]string{
		"home": homeTpl,
	}

	for k, v := range m {
		tpl, err := template.New(k).Parse(v)
		if err != nil {
			panic(err)
		}
		tpls[k] = tpl
	}

}
