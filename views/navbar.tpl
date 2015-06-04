{{define "navbar"}}

<nav class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <div class="navbar-header">
            <a class="navbar-brand" href="/">DOMAIN WHOIS</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
                <li{{if .isHome}}  class="active"{{end}}> <a href="/">首页</a> </li>
                <li{{if .IsCategory}}  class="active"{{end}}> <a href="/category">域名whois查询</a> </li>
                <li{{if .IsTopic}}  class="active"{{end}}> <a href="/topic">统计分析</a> </li>
                <li{{if .IsTopic}}  class="active"{{end}}> <a href="/topic">注册商查询</a> </li>
            </ul>
            <ul class="nav navbar-nav navbar-right">
                {{if .IsLogin}}
                <li><a href="/login?exit=true">退出</a></li>
                {{else}}
                <li><a href="/login">管理员登录</a></li>
                {{end}}
            </ul>
        </div>
    </div>
</nav>
{{end}}