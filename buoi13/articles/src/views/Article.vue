<template>
    <div class="article-page">
        <div class="banner">
            <div class="container">
                <h1>{{article.title}}</h1>

                <div class="article-meta">
                    <a href=""><img :src="author.image"/></a>
                    <div class="info">
                        <a href="" class="author">{{author.username}}</a>
                        <span class="date">{{article.createdAt}}</span>
                    </div>
                    <button class="btn btn-sm btn-outline-secondary">
                        <i class="ion-plus-round"></i>
                        &nbsp; Follow {{author.username}}
                    </button>
                    &nbsp;&nbsp;
                    <button class="btn btn-sm btn-outline-primary">
                        <i class="ion-heart"></i>
                        &nbsp; Favorite Post
                    </button>
                </div>
            </div>
        </div>

        <div class="container page">
            <div class="row article-content">
                <div class="col-md-12">
                    <p>
                        {{article.description}}
                    </p>
                    <p>
                       {{article.body}}
                    </p>
                </div>
            </div>

            <hr />

            <div class="article-actions">
                <div class="article-meta">
                    <a href="profile.html"
                        ><img :src="author.image"
                    /></a>
                    <div class="info">
                        <a href="" class="author">{{author.username}}</a>
                        <span class="date">{{article.creatAt}}</span>
                    </div>

                    <button class="btn btn-sm btn-outline-secondary">
                        <i class="ion-plus-round">Follow {{author.username}}</i>
                        &nbsp; 
                    </button>
                    &nbsp;
                    <button class="btn btn-sm btn-outline-primary">
                        <i class="ion-heart"></i>
                        &nbsp; Favorite Post <span class="counter"> </span>
                    </button>
                </div>
            </div>

            <div class="row">
                <div class="col-xs-12 col-md-8 offset-md-2">
                    <form class="card comment-form">
                        <div class="card-block">
                            <textarea
                                class="form-control"
                                placeholder="Write a comment..."
                                rows="3"
                            ></textarea>
                        </div>
                        <div class="card-footer">
                            <img
                                :src="user.image"
                                class="comment-author-img"
                            />
                            <button class="btn btn-sm btn-primary">
                                Post Comment
                            </button>
                        </div>
                    </form>

                    <div class="card"
                        v-for="comment in comments"
                        :key="comment.slug"
                    >
                        <div class="card-block">
                            <p class="card-text">
                               comment.body
                            </p>
                        </div>
                        <div class="card-footer">
                            <a href="" class="comment-author">
                                <img
                                    :src="comment.author.image"
                                    class="comment-author-img"
                                />
                            </a>
                            &nbsp;
                            <a href="" class="comment-author">{{comment.author.username}}</a>
                            <span class="date-posted">{{comment.createdAt}}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { mapState } from 'vuex';
import { FETCH_ARTICLE , FETCH_COMMENTS} from '../store/actions.type';
export default {
    data() {
        return {
        }
    },
    computed: {
        ...mapState({
            article: (state) => state.home.article ,
            comments: (state) => state.home.comments,
            user: state => state.auth.user,
            author: (state) => state.home.author
        }),
    },
    beforeMount() {
        let slug = this.$route.params.slug
        this.$store.dispatch(FETCH_ARTICLE, slug) 
        this.$store.dispatch(FETCH_COMMENTS, slug + "/comments")
    }
};
</script>
