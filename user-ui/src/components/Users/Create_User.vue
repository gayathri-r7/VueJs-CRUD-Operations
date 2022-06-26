<template>
<div class="container">


    <div class="row">
        <div class="col-6">
            <div class="input-group mb-3">
                <span class="input-group-text" id="basic-addon1">F</span>
                <input type="text" class="form-control" placeholder="First Name" aria-label="First Name" v-model="model.first_name" aria-describedby="basic-addon1">
            </div>
        </div>
        <div class="col-6">
            <div class="input-group mb-3">
                <span class="input-group-text" id="basic-addon2">L</span>
                <input type="text" class="form-control" placeholder="Last Name" aria-label="Last Name" v-model="model.last_name" aria-describedby="basic-addon2">
            </div>
        </div>

        <div class="col-6">
            <div class="input-group mb-3">
                <span class="input-group-text" id="basic-addon2"><i class='fas fa-briefcase'></i></span>
                <input type="text" class="form-control" placeholder="Job Title" aria-label="Job Title" v-model="model.job" aria-describedby="basic-addon2">
            </div>
        </div>
        <div class="col-6">
            <select class="form-select" v-model="model.gender" aria-label="Default select example">
                <option selected>Gender</option>
                <option value="female">Female</option>
                <option value="male">Male</option>
                <option value="other">Other</option>
            </select>
        </div>

        <div class="col-12">
            <div class="input-group mb-3">
                <input type="file" @change="previewFiles" class="form-control" id="inputGroupFile01">
                <label class="input-group-text" for="inputGroupFile01">Profile</label>
            </div>
        </div>

        <div class="col-4">
            <div class="input-group mb-3">
                <span class="input-group-text" id="basic-addon1"><i class='fab fa-linkedin fa-lg'></i></span>
                <input type="text" class="form-control" placeholder="LinkedIn" aria-label="LinkedIn" v-model="model.linkedin" aria-describedby="basic-addon1">
            </div>
        </div>
        <div class="col-4">
            <div class="input-group mb-3">
                <span class="input-group-text" id="basic-addon2"><i class="fab fa-twitter fa-lg"></i></span>
                <input type="text" class="form-control" placeholder="Twitter UserName" v-model="model.twitter" aria-label="Twitter UserName" aria-describedby="basic-addon2">
            </div>
        </div>
        <div class="col-4">
            <div class="input-group mb-3">
                <span class="input-group-text" id="basic-addon2"><i class="fab fa-skype fa-lg"></i></span>
                <input type="text" class="form-control" placeholder="Skype ID" aria-label="Skype ID" v-model="model.skype" aria-describedby="basic-addon2">
            </div>
        </div>

        <div class="col-12">
            <div class="input-group mb-3">
                <input type="text" class="form-control" placeholder="Recipient's Email" v-model="model.emailId" aria-label="Recipient's Email" aria-describedby="basic-addon2">
                <span class="input-group-text" id="basic-addon2">@example.com</span>
            </div>
        </div>

        <div v-if="operationType != 'Update'" class="col-6 mt-2">
            <div class="d-grid gap-2">
                <button class="btn btn-outline-primary" @click="reset()" type="button">Reset</button>
            </div>
        </div>
        <div v-if="operationType != 'Update'" class="col-6 mt-2">
            <div class="d-grid gap-2">
                <button class="btn btn-primary" @click="createUser()" type="button">Create</button>
            </div>
        </div>
    </div>
</div>
</template>

<script>

export default {
  name: 'CreateUser',
  props: ['operationType', 'user', 'entry'],
  data() {
      return {
        model: {
            first_name: "",
            last_name: "",
            job: "",
            gender: "female",
            profile_image: "",
            linkedin: "",
            twitter: "",
            skype: "",
            emailId: ""
        },
      }
    },
    watch: {
        'user': function(val) {
            if (this.operationType === 'Update') 
                this.model = val;
                this.$emit('saveUserData', val);
                this.$forceUpdate();
        },
        'model': function(val) {
            this.$emit('saveUserData', val);
        },
    },

    created() {
       
    },
    methods: {
        createUser(){
            var model = JSON.parse(JSON.stringify(this.model));
            this.$emit('createUser', model);
            this.reset();
        }, 

      //read user profile
      previewFiles(event) {
        this.model.profile_image = event.target.files[0]
        this.fileToDataURI(this.model.profile_image);
      },

      //File to data url
      fileToDataURI (imgFile) {
        var reader = new FileReader()
        reader.readAsDataURL(imgFile)
        reader.onload = () => {
          console.log('file to DataURI result:' + reader.result)
          this.model.profile_image = reader.result
        }
      },
    
     reset(){
        this.model.first_name = "";
        this.model.last_name = "";
        this.model.job = "";
        this.model.gender = "female";
        this.model.profile_image = "";
        this.model.linkedin = "";
        this.model.twitter = "";
        this.model.skype = "";
        this.model.emailId = "";
     }

    },
}
</script>

<style scoped>
@media (min-width: 1025px) {
    .h-custom {
        height: 100vh !important;
    }
}

</style>