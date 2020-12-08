<template>
  <div class="watch-rights">
     <!-- 密码设置 -->
    <div class="psw-set">
      <span class="psw-watch">密码观看</span>
      <el-switch  v-model="leftValue"  active-color="#dcdfe6"  inactive-color="#13ce66"  @change="handlePsw(leftValue)"></el-switch>
      <span class="psw-open">公开,所有人可见</span>
      <el-switch  v-model="rightValue"  active-color="#dcdfe6"  inactive-color="#13ce66"  @change="handleOpen(rightValue)"></el-switch>
    </div>
    <!-- 标签页 -->
    <el-tabs v-model="activeName">
      <el-tab-pane label="密码设置" name="first">
        <!-- 保存密码 -->
        <div class="player" >
          <div class="psw-hide">
            <span class="psw-note">往期/更新密码:</span> &nbsp;&nbsp;&nbsp;
            <el-input class="el-input " clearable  @input="handleClick"  size="small"
              placeholder="请输入6位数字密码" autofocus   maxlength="6"  v-model="roomData.password">
            </el-input>
          </div>
          <el-button type="primary" mini @click="pswSave">保存</el-button>
        </div>
      </el-tab-pane>
      <el-tab-pane label="白名单" name="second" style="width:90%">
        <div class="white-list">
          <el-button type="primary" @click="addListHandle">添加白名单</el-button>
        </div>
        <div class="list-row">
          <span>白名单</span>
          <span><el-button plain size="small">导出数据</el-button></span>
        </div>
        <!-- 白名单列表 -->
        <el-table v-loading="dataListLoading" border  ref="multipleTable"  tooltip-effect="dark"  style="width:100%" :data="userlists">
          <el-table-column type="selection" width="55"> </el-table-column>
          <el-table-column type=index  prop="" label="序号" width="120">
          </el-table-column>
          <el-table-column  prop="username"  label="用户名"  width="180"></el-table-column>
          <el-table-column  prop="phoneNumber"  label="手机号"  width="200" ></el-table-column>
          <el-table-column prop="operate" label="操作">
            <template slot-scope="scope">
              <el-button size="mini"  @click="showEditDialog(scope.$index,scope.row)">修改</el-button>
              <el-button size="mini" type="danger" @click="formDelete(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <!-- 分页 -->
        <el-pagination 
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="queryInfo.pagenum"
        :page-sizes="[5,10,15,20]"
        :page-size="queryInfo.pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total">
        </el-pagination>
      </el-tab-pane>
    </el-tabs>
    <!-- 添加白名单对话框 -->
    <addList v-if="addListVisible" :liveData="this.roomData" :centerDialog="isShow"  @sendVal="closeDialog" ref="addList"></addList>
    <!-- 修改对话框 -->
    <el-dialog   title="修改"  :visible.sync="editDialogVisible" width="50%"
      @keyup.enter.native="dataFormSubmitHandle()" center>
        <el-form :model="editForm"  :rules="rules"  ref="editForm"  label-width="100px"  class="demo-ruleForm">
          <el-form-item label="用户名" prop="username">
            <el-input v-model="editForm.username" placeholder="请输入用户名"></el-input>
          </el-form-item>
          <el-form-item label="用户手机号" prop="phoneNumber">
            <el-input  v-model="editForm.phoneNumber" placeholder="请输入手机号" ></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="editDialogVisible=false">取 消</el-button>
          <el-button type="primary" @click="dataFormSubmitHandle()">确 定</el-button>
        </span>
    </el-dialog>
  </div>
</template>

<script>
import mixinViewModule from '@/mixins/view-module';
import addList from "@/views/modules/avit/live/livehome/add-list";
import replayManage from "@/views/modules/avit/live/livehome/replay-manage"
import debounce from "lodash/debounce";
import $ from "jquery";
export default {
  mixins: [mixinViewModule],
  name: "watchRights",
  components: { addList },
  props:{
    roomData:Object
  },
  created(){
    this.getUserList()
    this.currentState()
    // console.log(this.roomData.viewPwd,'viewPwd')
    // console.log(this.roomData,'roomData')
  
  },
  data() {
      //   手机号验证
      var checkPhone = (rule, value, callback) => {
        const phoneReg = /^1[3|4|5|6|7|8][0-9]{9}$/;
        if (!value) {
          return callback(new Error("电话号码不能为空"));
        }
        setTimeout(() => {
          if (!Number.isInteger(+value)) {
            callback(new Error("请输入数字值"));
          } 
          else {
            if (phoneReg.test(value)) {
              callback();
            } 
            else {
              callback(new Error("电话号码格式不正确"));
            }
          }
        }, 100);
      };
      return {
        leftValue: true,
        rightValue:true,
          dataForm:{},
          mixinViewModuleOptions: {
            getDataListURL: '/live/roomwhitelink/page',
            getDataListIsPage: true,
            deleteURL: '/live/roomwhitelink',
            deleteIsBatch: true,
            exportURL: ''
          },
          editDialogVisible:false,
          userlists: [],
          total: 0,
          editForm:{},
          //获取用户列表的次数对象
          queryInfo: {
            query: '',
            pagenum: 1, //当前页数
            pagesize: 10,
          },
          activeName: "first",
          addListVisible: false,
          isShow:false,
          rules: {
            username: [
              { required: true, message: "请输入用户名", trigger: "blur" },
              { min: 1, max: 10, message: "长度在 1到 10个字符", trigger: "blur" },
            ],
            phoneNumber: [{ required: true, validator: checkPhone, trigger: "blur" }],
          },
      };
  },
  methods: {
    handleClick(e){
       let codeReg = new RegExp("[0-9]+"), //正则 英文+数字；
       len=e.length,
       str='';
       for(var i=0;i<len;i++){
          if(codeReg.test(e[i])){
            str+=e[i];
          }
        }
       this.roomData.password=str;
      },
    handlePsw(leftValue) {         //val为true时，不设置密码,禁用白名单，  为false时需要设置密码。开启白名单
      let switchData={
        id:this.roomData.id
      }
      if(leftValue){
        switchData.viewPwd=0
        this.$http.put("/live/liveroom/updateViewAuth",switchData).then(({ data: res }) => {
          if (res.code !== 0) {
            return this.$message.error(res.msg);
          }
          this.$message({
            message:'取消密码观看成功！',
            type: "success",
            duration: 1000,
          });
          this.roomData.password=""
        }).catch(() => {});
      }
      else{
        switchData.viewPwd=1
        this.$http.put("/live/liveroom/updateViewAuth",switchData).then(({ data: res }) => {
          if (res.code !== 0) {
            return this.$message.error(res.msg);
          }
          this.$message({
            message:'请设置密码！',
            type: "success",
            duration: 1000,
          });
          this.roomData.password=""
        }).catch(() => {});
      }
      
      },
    handleOpen(rightValue){
        let switchData={
          id:this.roomData.id
        }
        if(rightValue){
          switchData.viewAuthorized=0
          this.$http.put("/live/liveroom/updateViewAuth",switchData).then(({ data: res }) => {
            if (res.code !== 0) {
              return this.$message.error(res.msg);
            }
            this.$message({
              message:'取消所有人可见！',
              type: "success",
              duration: 1000,
            });
            this.roomData.password=""
          }).catch(() => {});
        }
        else{
          switchData.viewAuthorized=1
          this.$http.put("/live/liveroom/updateViewAuth",switchData).then(({ data: res }) => {
            if (res.code !== 0) {
              return this.$message.error(res.msg);
            }
            this.$message({
              message:'设置公开成功！',
              type: "success",
              duration: 1000,
            });
          }).catch(() => {});
        }
      },
    currentState(){
        //1.判断当前直播间用户是否设置了密码
        console.log(this.roomData.password,'初始状态')
        if(this.roomData.viewPwd===0){
          this.leftValue=true
        }
        else{
          this.leftValue=false 
        }
        if(this.roomData.viewAuthorized===0){
          this.rightValue=true
        }
        else{
          this.rightValue=false 
        }
      },
    addListHandle() {
        //todo:
        this.addListVisible = true;
        this.isShow=true;
      },
    closeDialog() {
        this.isShow=false;
      },
    showEditDialog(index,row){
        console.log(row)
        this.editDialogVisible=true
        this.editForm = JSON.parse(JSON.stringify(row))      //将浅拷贝带来的bug修改
      },
    pswSave(){
        let len=(this.roomData.password).length
        console.log(len,'len')
        if(len===0||len===6){
          let reset={
            password:this.roomData.password,
            id:this.roomData.id,
          }
          this.$http.put("/live/liveroom/updateViewAuth",reset).then(({ data: res }) => {
            if (res.code !== 0) {
                return this.$message.error(res.msg);
              }
            this.$message({
              message:'保存成功',
              type: "success",
              duration: 1000,
            });
          }).catch(() => {}); 
        }
        else{
          this.$message({
            message:'密码为6位数字！',
            type: "error",
            duration: 1000,
          });
          this.roomData.password=''
        }
      },
    // 获取用户列表
    getUserList(){  
        let  liveRoomId=this.roomData.id
        console.log(this.roomData.id,'id')
        this.$http.get("/live/roomwhitelink/page",{
          params: {
            liveRoomId:this.roomData.id,
            limit:this.queryInfo.pagesize,
            page:this.queryInfo.pagenum
          }
        }).then(({ data: res }) => {
          if (res.code !== 0) {
            return this.$message.error(res.msg);
          }
          console.log(res.data,'分页')
          this.userlists = res.data.list
          this.total = res.data.total
        }).catch(() => {})
      },
    //监听pagesize事件
    handleSizeChange(newSize) {
        console.log(newSize,'newSize')
        this.queryInfo.pagesize = newSize
        this.getUserList()
      },
    //页码值改变的事件
    handleCurrentChange(newPage) {
        this.queryInfo.pagenum = newPage
        console.log(newPage,'newPage')
        this.getUserList()
      },
    //表单提交
    dataFormSubmitHandle: debounce(
        function () {
          this.$refs["editForm"].validate((valid) => {
            if (!valid) return false;
            this.$http.put( "/live/roomwhitelink ", this.editForm).then(({ data: res }) => {
              if (res.code !== 0) {
                return this.$message.error(res.msg);
              }
              this.$message({
                message: this.$t("prompt.success"),
                type: "success",
                duration: 500,
              });
              this.getUserList()    //注意刷新代码放的位置
            }).catch(() => {});
            this.editDialogVisible=false
          });
        }, 1000, { leading: true, trailing: false }
      ),

    formDelete(row){
        this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(() => {
            this.$http.delete("/live/roomwhitelink", {'data':[row.id]} ).then(({ data: res }) => {
              if (res.code !== 0) {
                return this.$message.error(res.msg);
              }
              this.$message({
                message: this.$t("prompt.success"),
                type: "success",
                duration: 500,
              });
              this.getUserList() 
            }).catch(() => {});
        }).catch(() => {});
    }
  }
} 
  

</script>

<style scoped lang="scss">
.watch-rights {
  margin-left: 20px;
  .psw-set {
    margin-top: 30px;
    margin-bottom: 30px;
    .psw-watch{
       margin-right: 10px;
    }
    .psw-open {
      margin-left: 30px;
      margin-right: 20px;
    }
  }
  .player {
    .psw-hide {
      margin: 40px 60px;
      max-width:1024px;
      .el-input{
        width: 300px;
      }
      .psw-note {
        font-size: 16px;
      }
    }
  }
  .white-list {
    margin-top: 20px;
  }
  .list-row {
    display: flex;
    justify-content: space-between;
    margin-top: 20px;
    margin-bottom: 10px;
    span:nth-child(1) {
      margin: 8px;
    }
  }
}
</style>