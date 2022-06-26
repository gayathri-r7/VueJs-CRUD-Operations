<template>
  <div class="container">
    <div class="row">
        <div class="col">

            <b-table ref="table" stacked="md" :items="model" :fields="fields" class="table-style">   
                <template v-slot:cell(profile_image)="model">
                    <div>
                        <img class="img-fluid img-thumbnail" v-bind:src="model.item.profile_image" width="60" height="60"/>
                    </div>  
                </template>

                <template v-slot:cell(first_name)="model">
                    <div class="mt-3">{{model.item.first_name}}</div>
                </template>

                <template v-slot:cell(last_name)="model">
                    <div class="mt-3">{{model.item.last_name}}</div>
                </template>

                <template v-slot:cell(job)="model">
                    <div class="mt-3">{{model.item.job}}</div>
                </template>

                <template v-slot:cell(emailId)="model">
                    <div class="mt-3">{{model.item.emailId}}</div>
                </template>

                <template v-slot:cell(actions)="model">
                    <div class="mt-3">
                        <b-button class="mr-2" @click="getUser(model.item.id)" v-b-modal="'view-user-' + model.item.id">
                            <i class="fa fa-eye action-items-size" aria-hidden="true" title="View" style="color: DodgerBlue; cursor: pointer;"></i>
                        </b-button>
                        <b-button class="ml-2" @click="getUser(model.item.id)" v-b-modal="'edit-user-info' + model.item.id">
                            <i class="fa fa-pencil-square-o action-items-size" aria-hidden="true" title="Edit" style="color: Orange;cursor: pointer;"></i>
                        </b-button>
                        <b-button class="ml-3" v-b-modal="'delete-user-' + model.item.id">
                            <i class="fa fa-trash action-items-size" aria-hidden="true" title="Remove" style="color: red;cursor: pointer;"></i>
                        </b-button>
                    </div>

                    <b-modal :id="'edit-user-info' + model.item.id" title="Edit User" okTitle="Save" size="lg" @ok="saveUser(model.item.id)">
                        <EditUser operationType="Update" :user="getResponse" @saveUserData="userSavePayload($event)"></EditUser>
                    </b-modal>
                    <b-modal hide-footer :id="'view-user-' + model.item.id" size="sm">
                        <ViewUser :user="getResponse"></ViewUser>
                    </b-modal>
                    <b-modal :id="'delete-user-' + model.item.id" title="Delete User" okTitle="Delete" size="sm" @ok="deleteUser(model.item.id)">
                        <DeleteUser></DeleteUser>
                    </b-modal>

                </template>
            </b-table> 

            <div v-if="!model || model.length === 0" class="alert alert-primary alert-dismissible fade show text-center" role="alert">
                <strong>No users found!</strong> 
            </div>

        </div>      
    </div>
</div>

</template>

<script>

import EditUser from './Create_User.vue';
import ViewUser from './View_User.vue';
import DeleteUser from './Delete_User.vue';

import axios from "axios"

export default {
  components: {
    EditUser,
    ViewUser,
    DeleteUser
  },
  name: 'ListUser',
  props: ['userData'],
  data() {
      return {
        model: [],
        deleteResponse: {},
        createResponse: {},
        getResponse: {},
        entry: {},
        fields: [
            {
                key: 'profile_image',
                label: 'Profile'
            },
            {
                key: 'first_name',
                label: 'First Name'
            },
            {
                key: 'last_name',
                label: 'Last Name'
            },
            {
                key: 'job',
                label: 'Designation'
            },
            {
                key: 'emailId',
                label: 'Email ID'
            },
            {
                key: 'actions',
                label: 'Actions'
            },
        ],
      }
    },
    watch: {
        'userData': function(val) {
            console.log(val);
            this.createUser();
        }
    },
    created() {
        this.listUsers();
    },
    methods: {
        // Bootstarp toast message
        makeToast(variant = null, message) {
            this.$bvToast.toast(`${message}`, {
            title: `Variant ${variant || 'default'}`,
            variant: variant,
            solid: true
            })
        },
        // List Users
        listUsers(){
            var url = "/users";
            return axios({
                url: url,
                method: 'GET',
            }).then(result => {
                this.model = result.data;
                this.$emit('listUsers', this.model);
                return result;
            }).catch(error => {
                console.log(error);
                throw error;
            });
        },
        // Create User
        createUser(){
            var url = "/users";
            var model = JSON.parse(JSON.stringify(this.userData));

            return axios({
                url: url,
                method: 'POST',
                data: model
            }).then(result => {
                this.createResponse = result.data;
                this.listUsers();
                this.makeToast('success', "User created successfully.");
                return result;
            }).catch(error => {
                console.log(error);
                throw error;
            });
        }, 
        // Get User
        getUser(id){
            var url = "/users/" + id;
            return axios({
                url: url,
                method: 'GET',
            }).then(result => {
                this.getResponse = result.data;
                return result;
            }).catch(error => {
                console.log(error);
                throw error;
            });
        },
        // Save User
        saveUser(id){
            var url = "/users/" + id;
            this.entry.id = id;
            var model = JSON.parse(JSON.stringify(this.entry));

            return axios({
                url: url,
                method: 'PUT',
                data: model
            }).then(result => {
                console.log(result);
                this.createResponse = result.data;
                this.makeToast('success', 'User updated Successfully.');
                this.listUsers();
                return result;
            }).catch(error => {
                console.log(error);
                throw error;
            });
        }, 
        // Delete User
        deleteUser(id){
            var url = "/users/" + id;
            return axios({
                url: url,
                method: 'DELETE',
            }).then(result => {
                this.deleteResponse = result.data;
                this.listUsers();
                this.makeToast('success', 'User deleted successfully.');
                return result;
            }).catch(error => {
                console.log(error);
                throw error;
            });
        },
        // User Save Payload
        userSavePayload(data){
            this.entry = data;
        },

    },
    

}
</script>
<style scoped>
.action-items-size {
  font-size: 130%;
}
</style>
