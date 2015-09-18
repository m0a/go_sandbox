
Vue.config.debug = true;

// いくつかのコンポーネントを定義します
var FileList = Vue.extend({
    template: '#fileListTemplate',
    data: function(){
      console.log("inter data.")
      return {
        title: '',
        path: '',
        backpath:'',
        files: [],
        searchText: '',
      }
    },
    ready:function(){
      // console.log("inter compiled.");
      // console.log(this.$route.query.dir)
      if(this.$route.query.dir === undefined ) {
        this.getJson("./");
      } else {
        this.getJson(this.$route.query.dir);
      }


      // console.log("outerter compiled.");
    },
    methods:{
      nextLink:function(name) {
        console.log('inter nextLink');
        var nextDir = this.$data.path + '/' + name;
        var nextPath = {path: '/filelist?dir=' + nextDir };

         this.getJson(nextDir);
         this.$set('searchText','');
         this.$route.router.go(nextPath);

      },
      goBack:function() {
        console.log('inter backpath');
        var nextDir = this.$data.backpath;
        var nextPath = {path: '/filelist?dir=' + nextDir };
        this.$set('searchText','');
        this.getJson(nextDir);
        this.$route.router.go(nextPath);
      },
      playFile:function(name){
        console.log('inter playFile');
        var nextDir = this.$data.path + '/' + name;
        var nextPath = {path: '/playfile?file=' + nextDir };
        this.$route.router.go(nextPath);
      },
      getJson:function(dir){
        var that = this;
        $.ajax({
          type: 'GET',
          url: '/api/files?dir=' + dir,
          dataType: 'json',
          success: function(json) {
            that.$data.title = json[0].path;
            that.$data.path  = json[0].path;
            that.$data.backpath = json[0].path.replace(/\/[^\/]*$/,'');
            if (that.$data.backpath == '' ) {
              that.$data.backpath = '/';
            }
            that.$data.files = json[0].files;
          },
          data: null
        });
      }
    }
});

var PlayFile = Vue.extend({
    template: '#playFileTemplate',
    methods: {
      goBack:function() {
        var nextDir = this.$route.query.file.replace(/\/[^\/]*$/,'');
        var nextPath = {path: '/filelist?dir=' + nextDir };

         this.$route.router.go(nextPath);
      },
      srcFile:function() {
        var path = "/access" + this.$route.query.file
        console.log('srcFile='+path);
        return path;
      }
    }
});

// var App = Vue.extend({});

Vue.filter('backpath',function(str){
  var s = str.replace(/\/[^\/]*$/,'');
  console.log(s);
  return s;
});

Vue.filter('short',function(str){
  var s = str.replace(/.*(.{30})$/,'$1');
  console.log(s);
  return s;
});

$(function(){



  var App =  Vue.extend({});

  // router インスタンスを作成。
  // ここでは追加的なオプションで渡すことができますが、今はシンプルに保っています
  var router = new VueRouter()

  // いくつかの routes を定義します
  // route 毎、コンポーネントにマップが必要です
  // "component" は 事実上コンポーネントコンストラクタは Vue.extend() 経由で作成されるか、
  // または適切なコンポーネントオプションオブジェクトでできます
  // nested routes については後で話します
  router.map({
      '/': {
        component: FileList
      },
      '/filelist': {
          component: FileList
      },
      '/playfile': {
          component: PlayFile
      }
  })

  // 今 アプリケーションを開始することが出来ます！
  // router は App のインスタンスを作成し、
  // そして #app セレクタでマッチングした要素にマウントします
  router.start(App,'#app')
});
