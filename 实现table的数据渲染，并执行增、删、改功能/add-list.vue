<template>
  <!-- 白名单列表 -->
  <!-- .sync触发的是父组件的update事件 -->
  <el-dialog  
    title="添加白名单"  :visible.sync="dialogVisible" width="50%"  center>
    <el-form :model="dataForm"  :rules="rules"  ref="dataForm"  label-width="100px"  class="demo-ruleForm">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="dataForm.username" placeholder="请输入用户名"></el-input>
      </el-form-item>
      <el-form-item label="用户手机号" prop="phoneNumber">
        <el-input  v-model="dataForm.phoneNumber" placeholder="请输入手机号" ></el-input>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button @click="dialogVisible='false'">取 消</el-button>
      <el-button type="primary" @click="updateList">确 定</el-button>
    </span>
  </el-dialog>
</template>

<script>
import debounce from "lodash/debounce";
export default {
  name:'addList',
   props:{
     liveData: Object,
    centerDialog:Boolean
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
        } else {
          if (phoneReg.test(value)) {
            callback();
          } else {
            callback(new Error("电话号码格式不正确"));
          }
        }
      }, 100);
    };
    return {
    
      formLabelWidth: "120px",
      dataForm: {
        username: "",
        phoneNumber: "",
        liveRoomId: this.liveData.id,
      },
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
   updateList: debounce(
      function () {
        this.$refs["dataForm"].validate((valid) => {
          if (!valid) return false;
          this.$http.post( "/live/roomwhitelink", this.dataForm).then(({ data: res }) => {
            if (res.code !== 0) {
              return this.$message.error(res.msg);
            }
            this.$message({
              message: this.$t("prompt.success"),
              type: "success",
              duration: 500,
              onClose: () => {
                 this.close();
                this.$parent.getUserList()
                console.log("刷新成功！")
               
              },
            });
          }).catch(() => {});
        });
      }, 1000, { leading: true, trailing: false }
    ),
    close() {
      this.$emit('sendVal')
    }
  },
  computed:{
     dialogVisible: {
        get(){
          return this.centerDialog  //获取父组件的值
        },
        set(val){
        this.$emit('sendVal',val)     //将子组件改变的值传递给父组件
    },

      },
  },
  created(){

  }
};
</script>

<style scoped>
</style>
