<template>
<div class="container">
    <div class="row">
        <div class="card mb-4">
          <div class="card-body text-center">
            <img v-bind:src="user.profile_image" alt="avatar" :class="user.gender === 'male' ? 'rounded-circle img-fluid img-male' : 'rounded-circle img-fluid img-female'" style="width: 150px;">
            <h5 class="my-3">{{user.first_name}} {{user.last_name}}</h5>
            <p class="text-muted">{{user.job}}</p>

            <div class="mb-4 pb-2 mt-3">
              <button type="button" class="btn btn-outline-primary btn-floating">
                <i class="fa fa-linkedin"></i>
              </button>
              <button type="button" class="btn btn-outline-primary ms-2">
                <i class="fab fa-twitter fa-lg"></i>
              </button>
              <button type="button" class="btn btn-outline-primary btn-floating ms-2">
                <i class="fab fa-skype fa-lg"></i>
              </button>
            </div>
          
            <div class="d-flex justify-content-center mt-2 mb-2">
            <button type="button" class="btn btn-outline-primary">Message</button>
              <button type="button" class="btn btn-primary ms-2">Follow</button>
              
            </div>
          </div>
        </div>
    </div>
</div>
</template>

<script>
export default {
  name: 'ViewUser',
  props: ['user'],
}
</script>

<style scoped>
.img-male {
  border: 1px solid;
  color: Orange;
}
.img-female {
  border: 2px solid;
  color: pink;
}
</style>