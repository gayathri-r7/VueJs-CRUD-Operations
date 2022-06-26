<template>
<div>
    
    <b-row md="12">
        <b-col md="12">
            <div class="card">
                <div class="card-body">
                    <h5 class="card-title"><i class='fas fa-users user-icon-style'></i><span class="user-title-style">  USERS</span></h5>
                    <p v-if="usersCount != 0" class="text-center"><span class="badge rounded-pill bg-success user-count-style mt-3">{{usersCount}}</span></p>
                    <p v-else class="text-center"><span class="badge rounded-pill bg-danger user-count-style mt-3">0</span></p>
                </div>
            </div>
        </b-col>
    </b-row>

</div>
</template>

<script>
export default {
  name: 'DisplayUser',
  props: ['usersCount'],
}
</script>
<style scoped>
.user-icon-style { 
    font-size: 50px;
    color: orange;
}

.user-title-style { 
    font-size: 32px;
    color: orange;
}
.user-count-style {
    font-size: 25px;
}
</style>
