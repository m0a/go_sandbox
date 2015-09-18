
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
        // backupSearchText:{},
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
        // this.$set('searchText','');
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


Vue.filter('backpath',function(str){
  var s = str.replace(/\/[^\/]*$/,'');
  console.log(s);
  return s;
});

Vue.filter('short',function(str){
  var s = str.replace(/.*(.{45})$/,'$1');
  console.log(s);
  return s;
});

$(function(){

  //Router設定
  var App =  Vue.extend({});
  var router = new VueRouter()
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
  });

  router.start(App,'#app')

});
