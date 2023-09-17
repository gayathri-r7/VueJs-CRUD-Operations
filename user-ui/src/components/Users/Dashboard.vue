<template>
  <div class="container">
    
    <div class="row mt-5">
      <div class="col-9">
        <CreateUser @createUser="userCreatePayload($event)" />
      </div>
      <div class="col-3">
        <DisplayBoard :usersCount="usersCount" />
      </div>
    </div>

    <div class="row mt-5">
      <div class="col-12 mt-3">
        <Users @listUsers="Users($event)" :userData="data" />
      </div>
    </div>

</div>
</template>

<script>

import CreateUser from './Create_User.vue'
import DisplayBoard from './Display_Users.vue'
import Users from './Users.vue'

export default {
  components: {
    CreateUser,
    DisplayBoard,
    Users
  },
  name: 'DashboardUser',
  data() {
      return {
        data: {},
        usersCount: 0,
      }
  },
  watch: {
    
  },
  created() {
    
    
  },
  methods: {
    userCreatePayload(data){
      this.data = data;
    },
    Users(data){
      if (data && data.length != 0) this.usersCount = data.length;
      else this.usersCount = 0;
    }
  }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
