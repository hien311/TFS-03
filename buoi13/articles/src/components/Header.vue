<template>
  <nav class="navbar navbar-light">
    <div class="container">
      <a class="navbar-brand" href="/">conduit</a>
      <ul class="nav navbar-nav pull-xs-right" v-if="!isAuthenticated">
        <li class="nav-item">
          <router-link
            class="nav-link"
            active-class="active"
            :to="{ name: 'home' }"
          >
            Home
          </router-link>
        </li>
        <li class="nav-item">
          <router-link
            class="nav-link"
            active-class="active"
            :to="{ name: 'login' }"
          >
            Login
          </router-link>
        </li>
        <li class="nav-item">
          <router-link
            class="nav-link"
            active-class="active"
            :to="{ name: 'register' }"
          >
            Register
          </router-link>
        </li>
      </ul>
      <ul class="nav navbar-nav pull-xs-right" v-else>
        <li class="nav-item">
          <!-- Add "active" class when you're on that page" -->
          <router-link
            class="nav-link"
            active-class="active"
            :to="{ name: 'home' }"
          >
            Home
          </router-link>
        </li>
        <li class="nav-item">
             <router-link 
            class="nav-link"
            active-class="active"
            :to="{ name : 'new-article'}"
          >
            New Article
          </router-link>
        </li>
        <li class="nav-item">
          <router-link 
            class="nav-link"
            active-class="active"
            :to="{ name : 'setting'}"
          >
            Setting
          </router-link>
        </li>
        <li class="nav-item">
          
          <router-link 
            class="nav-link"
            active-class="active"
            :to="{ name : 'user'}"
          >
            {{userName}}
          </router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "Header",
  computed: {
    ...mapState({
      isAuthenticated: (state) => state.auth.isAuthenticated,
    }),
    userName() {
      return this.$store.state.auth.user.username
    }
  },
};
</script>
