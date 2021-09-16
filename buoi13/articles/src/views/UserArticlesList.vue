<template>
    <div id="user-list">
        <div class="article-preview" v-if="userList==null">
            <div class="article-meta">
                No article are here ...yet
            </div>
        </div>
        <div class="article-preview"
            v-else
            v-for="article in userList"
            :key="article.slug"
        >
            <div class="article-meta">
                <a href=""><img :src="article.author.image"/></a>
                <div class="info">
                    <a href="" class="author">{{article.author.username}}</a>
                    <span class="date">{{article.createdAt}}</span>
                </div>
                <button class="btn btn-outline-primary btn-sm pull-xs-right">
                    <i class="ion-heart"></i> 29
                </button>
            </div>
            <a href="" class="preview-link">
                <h1>{{article.title}}</h1>
                <p>{{article.description}}</p>
                <span>Read more... </span>
            </a>
        </div>
    </div>
</template>

<script>
import { FETCH_USER_ARTICLES } from '../store/actions.type'
import { mapState } from 'vuex'
export default {
    computed: {
        ...mapState({
            userList: (state) => state.home.userList,
            user: (state) => state.auth.user
        }),
        type() {
            return this.$route.params.type
        }
    },
    beforeMount() {
        this.$store.dispatch(FETCH_USER_ARTICLES, {
            [this.type]: this.user.username,
            offset: 0,
            limit: 10
            })
    },
    watch: {
        type: function() {
            this.$store.dispatch(FETCH_USER_ARTICLES, {
            [this.type]: this.user.username,
            offset: 0,
            limit: 10
            })
        }
    }
}
</script>