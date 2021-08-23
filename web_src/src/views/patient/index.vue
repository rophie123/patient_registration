<template>
  <div class="dashboard-container">
    <h3>Patient List</h3>
    <el-table
      :data="patients"
      style="width: 100%">
      <el-table-column type="index" width="50">
      </el-table-column>
      <el-table-column
        prop="Name"
        label="Name">
      </el-table-column>
      <el-table-column
        prop="Birth"
        label="Birth">
      </el-table-column>
      <el-table-column
        prop="Phone"
        label="Phone">
      </el-table-column>
      <el-table-column
        prop="Email"
        label="Email">
      </el-table-column>
      <el-table-column
        prop="Address"
        label="Address">
      </el-table-column>
      <el-table-column
        label="Photo">
        <template slot-scope="scope">
          <img :src="url+scope.row.Photo" width="80">
        </template>
      </el-table-column>
      <el-table-column
        prop="AppointmentTime"
        label="Appointment Time">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {mapGetters} from 'vuex'
import {userList} from "@/api/user";

export default {
  name: 'Patients',
  computed: {
    ...mapGetters([
      'name',
      'roles'
    ])
  }, data() {
    return {
      patients: [],
      url:''
    }
  },
  created() {
    this.url=process.env.VUE_APP_BASE_API+'/upload/'
    this.loadPatients()
  },
  methods: {
    loadPatients() {
      userList().then(res => {
        this.patients = res.data
        console.log(res.data)
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
&-container {
  margin: 30px;
}

&-text {
  font-size: 30px;
  line-height: 46px;
}
}
</style>
