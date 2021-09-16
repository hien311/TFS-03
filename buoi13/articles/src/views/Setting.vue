<template>
   <div class="settings-page">
  <div class="container page">
    <div class="row">

      <div class="col-md-6 offset-md-3 col-xs-12">
        <h1 class="text-xs-center">Your Settings</h1>

        <form>
          <fieldset>
              <fieldset class="form-group">
                <input class="form-control" type="text" placeholder="URL of profile picture" v-model="image">
              </fieldset>
              <fieldset class="form-group">
                <input class="form-control form-control-lg" type="text" :placeholder="user.username" v-model="username">
              </fieldset>
              <fieldset class="form-group">
                <textarea class="form-control form-control-lg" rows="8" :placeholder=" user.bio == null ? 'Your bio' : user.bio " v-model="bio"></textarea>
              </fieldset>
              <fieldset class="form-group">
                <input class="form-control form-control-lg" type="text" :placeholder="user.email" v-model="email">
              </fieldset>
              <fieldset class="form-group">
                <input class="form-control form-control-lg" type="password" placeholder="Password" v-model="password">
              </fieldset>
              <button class="btn btn-lg btn-primary pull-xs-right" @click.prevent="submit">
                Update Settings
              </button>
              <button class="btn btn-lg btn-danger pull-xs-left" @click="loggout">
                Loggout
              </button>
          </fieldset>
        </form>
      </div>

    </div>
  </div>
</div>
</template>

<script>
import { UPDATE_PROFILE, LOGOUT } from '../store/actions.type'
import { mapState } from 'vuex'
export default {
  data() {
    return {
      email: "",
      token: "",
      username: "",
      bio: "",
      image: "",
      password: "",
    }
  },
  computed: {
    ...mapState({
      user: (state) => state.auth.user
    })
  },
  methods: {
    async submit() {
      const payload = {
        email: this.email === "" ? this.user.email : this.email,
        token: this.token === "" ? this.user.token : this.token,
        username: this.username === "" ? this.user.username : this.username,
        bio: this.bio === null ? this.user.bio : this.bio,
        image: this.image === null ? this.user.image : this.image,
      }
      if (await this.$store.dispatch(UPDATE_PROFILE, payload )) {
        this.$router.push({ name : 'home' })
      }
    },
    async loggout() {
      if ( await this.$store.dispatch(LOGOUT) ) {
        this.$router.push({name: 'login'})
      }
    }
  }
}
</script>