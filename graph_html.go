package main

var GraphHtmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>iview example</title>
    <link rel="stylesheet" type="text/css" href="http://unpkg.com/view-design/dist/styles/iview.css">
    <script type="text/javascript" src="http://vuejs.org/js/vue.min.js"></script>
    <script type="text/javascript" src="http://unpkg.com/view-design/dist/iview.min.js"></script>
</head>
<body>
<div id="app" style="margin: auto auto auto 10px;">
    <i-input v-model="keyword" placeholder="请输入..." style="width: 600px; margin: 5px auto 5px auto;"></i-input>
    <i-button @click="doSearch">搜索</i-button>
    <i-button @click="switchExpandStatus">{{ title }}</i-button>
    <p style="color:red; margin: 5px auto 5px auto;">输入关键字，按[Enter]键实现搜索</p>
    <Tree :data="showData" :render="renderContent"></Tree>
</div>
<script>
    new Vue({
        el: '#app',
        computed: {
            showData: function () {
                if (!this.useSearchData || !this.keyword || this.keyword.length < 1) {
                    return this.baseData
                } else {
                    return this.searchData
                }
            }
        },
        data: {
            expanded: true,
            title: '全部收起',
            keyword: "",
            useSearchData: false,
            baseData: [],
            searchData: [],
            pathNodes: [],
            nodeMap: {}
        },
        created: function () {
            this.baseData =${graphJsonData}
            // 构造pathNodes

            let list = []
            this.baseData.forEach(function (item) {
                list.push(item)
            })

            for (let i = 0; i < list.length; i++) {
                let node = list[i]
                this.nodeMap[node.title] = node
                let children = node.children
                if (children && children.length > 0) {
                    children.forEach(function (item) {
                        list.push(item)
                    })
                }
            }

            for (let id in this.nodeMap) {
                let paths = [id]
                let cnode = this.nodeMap[id]
                while (true) {
                    if (!cnode.pid || cnode.pid.length < 1) {
                        break
                    }
                    cnode = this.nodeMap[cnode.pid]
                    paths.push(cnode.title)
                }
                this.pathNodes.push(paths)
            }
        },
        destroyed () {
            // 销毁enter事件
            this.enterKeyupDestroyed()
        },
        mounted () {
            // 绑定enter事件
            this.enterKeyup()
        },
        methods: {
            enterKeyupDestroyed () {
                document.removeEventListener('keyup', this.enterKey)
            },
            enterKeyup () {
                document.addEventListener('keyup', this.enterKey)
            },
            enterKey (event) {
                const code = event.keyCode
                    ? event.keyCode
                    : event.which
                        ? event.which
                        : event.charCode
                // eslint-disable-next-line eqeqeq
                if (code === 13) {
                    this.doSearch()
                }
            },
            renderContent (h, { root, node, data }) {
                if (!this.useSearchData || !this.keyword || this.keyword.length < 1) {
                    return h('span', {}, data.title)
                }
                let idx = data.title.indexOf(this.keyword)
                if (idx < 0) {
                    return h('span', {}, data.title)
                }
                let array = data.title.split(this.keyword)
                let kspan = h('span', {style: {color: 'red' } }, this.keyword)
                let spans = []
                for (let i=0; i<array.length; i++) {
                    spans.push(h('span', {}, array[i]))
                    if (i !== array.length - 1) {
                        spans.push(kspan)
                    }
                }
                return h('span', {
                    style: {
                        display: 'inline-block',
                        width: '100%'
                    }
                }, spans);
            },
            switchExpandStatus: function() {
                let list = this.baseData
                if (this.useSearchData) {
                    list = this.searchData
                }
                let iterList = []
                list.forEach(function(item) {
                    iterList.push(item)
                })
                for (let i=0;i<iterList.length; ++i ) {
                    iterList[i].expand = !iterList[i].expand
                    if (iterList[i].children) {
                        iterList[i].children.forEach(function(item){
                            iterList.push(item)
                        })
                    }
                }

                this.expanded = !this.expanded
                if (this.expanded) {
                    this.title = "全部收起"
                } else {
                    this.title = "全部展开"
                }
            },
            doSearch: function () {
                if (this.keyword.length < 1) {
                    this.useSearchData = false
                    return
                }
                this.useSearchData = true

                let keyword = this.keyword
                let nodes = []
                let exists = {}
                this.pathNodes.forEach(function (paths) {
                    for (let i = 0; i < paths.length; ++i) {
                        if (paths[i].indexOf(keyword) >= 0) {
                            // 找到了，构造 nodes
                            let pNode = null
                            let tNode = null
                            let path = ''
                            for (let j = paths.length - 1; j >= i; j--) {
                                path = path + paths[j]
                                let node = {title: paths[j], expand: true}
                                if (pNode != null) {
                                    node.pid = pNode.title
                                    pNode.children = [node]
                                } else {
                                    tNode = node
                                }
                                pNode = node
                            }
                            if (!exists[path]) {
                                nodes.push(tNode)
                                exists[path] = true
                            }
                            break
                        }
                    }
                })
                this.searchData = nodes
            }
        }
    })
</script>
</body>
</html>
`
