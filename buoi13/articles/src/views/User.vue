<template>
    <div class="profile-page">
        <div class="user-info">
            <div class="container">
                <div class="row">
                    <div class="col-xs-12 col-md-10 offset-md-1">
                        <img :src="user.image" class="user-img" />
                        <h4>{{ user.username }}</h4>
                        <p>
                            {{ user.bio }}
                        </p>
                        <button
                            class="btn btn-sm btn-outline-secondary action-btn"
                        >
                            <i class="ion-plus-round"></i>
                            &nbsp; Follow {{ user.username }}
                        </button>
                    </div>
                </div>
            </div>
        </div>

        <div class="container">
            <div class="row">
                <div class="col-xs-12 col-md-10 offset-md-1">
                    <div class="articles-toggle">
                        <ul class="nav nav-pills outline-active">
                            <li class="nav-item">
                                <router-link
                                    class="nav-link"
                                    active-class="active"
                                    :to="{ name: 'user-articles-list', params:{ type:'author'}}"
                                >
                                    My Articles
                                </router-link>
                            </li>
                            <li class="nav-item">
                                <router-link
                                    class="nav-link"
                                    active-class="active"
                                    :to="{ name: 'user-articles-list', params:{ type:'favorited'}}"
                                >
                                    Favorited Articles
                                </router-link>
                            </li>
                        </ul>
                    </div>
                <router-view />
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { mapState } from "vuex";
export default {
    computed: {
        ...mapState({
            user: (state) => state.auth.user,
        }),
    },
};
</script>
