<template>
  <div class="container">
    <el-form ref="patientForm" :model="patientForm" :rules="rules" auto-complete="on" class="form"
             label-position="top">

      <div class="title-container">
        <h3 class="title">Patient Registration</h3>
      </div>

      <el-form-item label="Name" prop="name">
        <el-input
          ref="name"
          v-model="patientForm.name"/>
      </el-form-item>
      <el-form-item label="Date of birth" prop="birth">
        <el-date-picker style="width: 100%"
                        v-model="patientForm.birth"
                        ref="birth"
                        type="date"
                        :picker-options="picker_birth_option">
        </el-date-picker>
      </el-form-item>
      <el-form-item label="Phone" prop="phone">
        <el-input
          ref="phone"
          v-model="patientForm.phone"/>
      </el-form-item>
      <el-form-item label="Email" prop="email">
        <el-input
          ref="email"
          v-model="patientForm.email"/>
      </el-form-item>
      <el-form-item label="Address" prop="address">
        <el-input
          ref="address"
          v-model="patientForm.address"/>
      </el-form-item>
      <el-form-item label="Photo" prop="photo">
        <el-upload
          class="avatar-uploader"
          :action="action"
          :show-file-list="false"
          :on-success="handleUploadSuccess"
          :before-upload="beforeUpload">
          <img v-if="imageUrl" :src="imageUrl" class="avatar">
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
      </el-form-item>
      <el-form-item label="Appointment Time" prop="appointment_time">
        <el-date-picker style="width: 100%"
                        v-model="patientForm.appointment_time"
                        ref="appointment_time"
                        type="datetime"
                        :picker-options="picker_appointment_time_option">
        </el-date-picker>
      </el-form-item>
      <el-button :loading="loading" type="primary" style="margin-top: 30px" @click.native.prevent="handleClick">Save
      </el-button>
    </el-form>
  </div>
</template>

<script>
import {isEmpty, validEmail, validPhone} from "@/utils/validate";
import {Reg} from "@/api/user";

export default {
  name: "AddPatient",
  data() {
    const validateName = (rule, value, callback) => {
      if (isEmpty(value)) {
        callback(new Error('Please enter the your name'))
      } else {
        callback()
      }
    }
    const validateBirth = (rule, value, callback) => {
      if (isEmpty(value)) {
        callback(new Error('Please enter the your date of birth'))
      } else {
        callback()
      }
    }
    const validatePhone = (rule, value, callback) => {
      if (!validPhone(value)) {
        callback(new Error('Please enter the your phone,example 13811111111'))
      } else {
        callback()
      }
    }
    const validateEmail = (rule, value, callback) => {
      if (!validEmail(value)) {
        callback(new Error('Please enter the your email'))
      } else {
        callback()
      }
    }
    const validateAddress = (rule, value, callback) => {
      if (isEmpty(value)) {
        callback(new Error('Please enter the your address'))
      } else {
        callback()
      }
    }
    const validatePhoto = (rule, value, callback) => {
      if (this.patientForm.photo === '') {
        callback(new Error('Please upload the photo'))
      } else {
        callback()
      }
    }
    const validateAppointmentTime = (rule, value, callback) => {
      if (isEmpty(value)) {
        callback(new Error('Please enter the your appointment time'))
      } else {
        callback()
      }
    }
    return {
      picker_birth_option:{
        disabledDate:(time)=> {
          return time.getTime() > Date.now()-1 * 24 * 3600 * 1000
        }
      },
      picker_appointment_time_option:{
        disabledDate:(time)=> {
          return time.getTime() < Date.now()-1 * 24 * 3600 * 1000
        }
      },
      patientForm: {
        name: '',
        birth: '',
        phone: '',
        email: '',
        address: '',
        photo: '',
        appointment_time: ''
      },
      rules: {
        name: [{required: true, trigger: 'blur', validator: validateName}],
        birth: [{required: true, trigger: 'blur', validator: validateBirth}],
        phone: [{required: true, trigger: 'blur', validator: validatePhone}],
        email: [{required: true, trigger: 'blur', validator: validateEmail}],
        address: [{required: true, trigger: 'blur', validator: validateAddress}],
        photo: [{required: true, trigger: 'blur', validator: validatePhoto}],
        appointment_time: [{required: true, trigger: 'blur', validator: validateAppointmentTime}],
      },
      loading: false,
      imageUrl: '',
      action:''
    }
  },
  created() {
    this.action=process.env.VUE_APP_BASE_API+'/photo/upload'
  },
  methods: {
    handleClick() {
      this.$refs.patientForm.validate(valid => {
        if (valid) {
          this.loading = true
          Reg(this.patientForm).then(res => {
            this.$alert(res.message)
          }).finally(() => {
            this.loading = false
          })
        } else {
          return false
        }
      })
    },
    handleUploadSuccess(res, file) {
      this.imageUrl = URL.createObjectURL(file.raw);
      this.patientForm.photo = file.name
    },
    beforeUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('You can only upload pictures in JPG format');
      }
      if (!isLt2M) {
        this.$message.error('The size of the uploaded picture cannot exceed 2MB');
      }
      return isJPG && isLt2M;
    }
  }
}
</script>

<style lang="scss" scoped>

.container {
  min-height: 100%;
  width: 100%;
  overflow: hidden;

.form {
  position: relative;
  width: 520px;
  max-width: 100%;
  padding: 35px 35px 0;
  margin: 0 auto;
  overflow: hidden;
}

.tips {
  font-size: 14px;
  color: #fff;
  margin-bottom: 10px;
}


.title-container {
  position: relative;
}

.title {
  font-size: 26px;
  margin: 0px auto 40px auto;
  text-align: center;
  font-weight: bold;
}

}
</style>
<style>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 100px;
  height: 100px;
  line-height: 100px;
  text-align: center;
}

.avatar {
  width: 100px;
  height: 100px;
  display: block;
}
</style>
