<template>
<div id="app" class="container">
    <div>
        <title><img src="assets/img/mockturnal.jpg" alt="">GoScreen</title>
    </div>
    <div class="row">
        <div class="col-md-6 offset-md-3 py-5">
            <h1>Generate a thumbnail of a website</h1>

            <form v-on:submit.prevent="makeWebsiteThumbnail">
                <div class="form-group">
                    <input v-model="websiteUrl" type="text" id="website-input" placeholder="Enter a website" class="form-control">
                </div>
                <div class="form-group">
                    <button class="btn btn-primary">Generate!</button>
                </div>
                <div>
                    <img class="img" :src="thumbnailUrl">
                </div>
            </form>
        </div>
    </div>
</div>
</template>

<script>
import axios from 'axios';

export default {
    name: 'App',
    components: {},
    data() {
        return {
            websiteUrl: '',
            thumbnailUrl: ''
        }
    },
    methods: {
        makeWebsiteThumbnail() {
            axios.post("http://localhost:3000/api/thumbnail", {
                    url: this.websiteUrl,
                })
                .then((response) => {
                    this.thumbnailUrl = response.data.screenshot;
                })
                .catch((error) => {
                    window.alert(`The API returned an error: ${error}`);
                })
        }
    },
}
</script>

<style>
#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
}

.img {
    width: 500px;
}
</style>
