<template>
  <div class="regContainer">
    <div class="regBox">
      <div class="title">
        <h1>注册</h1>
      </div>
      <div class="regForm">
        <el-form ref="form" :rules="rules" :model="form" class="loginForm" size="middle">
          <el-form-item prop="username">
            <el-input v-model="form.username" prefix-icon="el-icon-mobile-phone" placeholder="请输入用户名"></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input v-model="form.password" type="password" prefix-icon="el-icon-lock" placeholder="请输入密码"></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input v-model="form.rpassword" type="password" prefix-icon="el-icon-lock" placeholder="请再次输入密码"></el-input>
          </el-form-item>
          <el-form-item prop="captcha_code">
            <div class="captcha">
              <el-input v-model="form.captcha_code" placeholder="图片验证码" prefix-icon="el-icon-document"></el-input>
              <img @click="getCaptchaImg" class="captcha__img" :src="captchaUri" alt="" />
            </div>
          </el-form-item>
        </el-form>
      </div>
      <div class="btn">
        <el-button class="loginBtn" type="primary" @click="toLogin('form')">登录</el-button>
        <el-button class="regBtn" type="primary" @click="submitForm('form')">注册</el-button>
      </div>
    </div>
  </div>
</template>

<script>
import { imageCaptchaAPI, registerAPI } from "@/api/sign-in/login";
export default ({
  name: '',
  data() {
    return {
      captchaUri: "",
      form: {
        username: 'admin',
        rpassword: 'wulinlin',
        captcha_id: '',
        captcha_code: '',
        password: 'wulinlin'
      },
      rules: {
        username: [
          {
            required: true,
            message: '请输入用户名',
            trigger: 'blur'
          },
          {
            min: 3,
            max: 15,
            message: '用户名长度不合法',
            trigger: 'blur'
          }
        ],
        password: [
          {
            required: true,
            message: '请输入密码',
            trigger: 'blur'
          },
          {
            min: 6,
            max: 15,
            message: '密码长度不合法',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  created() {
    this.getCaptchaImg();
  },
  methods: {
    toLogin() {
      this.$router.push('/login')
    },
    submitForm(form) {
      this.$refs[form].validate(async valid => {
        if (!valid) return this.$message.error('非法输入数据，请重新输入')
        registerAPI(this.form).then((res) => {
          console.log(res)
        }).catch((e) => {
          this.$message({
            message: e.msg,
            type: "error",
          });
        });
      })
    },
    getCaptchaImg() {
      imageCaptchaAPI()
        .then((res) => {
          this.form.captcha_id = res.captchaId;
          this.captchaUri = res.picPath;
        })
        .catch((e) => {
          this.$message({
            message: e.msg,
            type: "error",
          });
        });
    },
  }
})
</script>

<style scoped>
.regContainer {
  height: 100%;
  /* background-color: #282c34; */
}
.regBox {
  width: 450px;
  /* height: 450px; */
  position: absolute;
  top: 45%;
  left: 50%;
  transform: translate(-50%, -60%);
  /*border-radius: 20px;*/
  background-color: #ffffff; /* 容器背景颜色 */
  border: 1px solid #dddddd; /* 边框 */
  border-radius: 10px; /* 圆角 */
  box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.1); /* 阴影 */
  padding: 20px; /* 内边距 */
  margin: 20px; /* 外边距 */
}
.title {
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}

.captcha {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.captcha__img {
  margin-left: 20px;
  width: 150px;
  height: 40px;
  border-bottom: 1px solid #dbdbdb;
}
.btn {
  display: flex;
  justify-content: center;
  margin-top: 20px;
  margin-bottom: 30px;
}
.regBtn {
  color: #ffffff;
  background: linear-gradient(to right, #ecf0f1, #3498db);
  margin-left: 30px;
}
</style>