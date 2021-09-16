<template>
  <main id="global-feed">
      <div
            class="article-preview"
            v-for="article in articles"
            :key="article.slug"
          >
            <div class="article-meta">
              <a href="profile.html"
                ><img :src="article.author.image"
              /></a>
              <div class="info">
                <a href="" class="author">{{ article.author.username }}</a>
                <span class="date">{{ article.createdAt }}</span>
              </div>
              <button class="btn btn-outline-primary btn-sm pull-xs-right">
                <i class="ion-heart"></i> {{ article.favoritesCount }}
              </button>
            </div>
            <a href="" class="preview-link">
              <h1>{{ article.title }}</h1>
              <p>{{ article.description }}</p>
              <span @click="readDetail(article.slug)">Read more...</span>
            </a>
          </div>
  </main>
</template>

<script>
import { mapState} from 'vuex'
import {  FETCH_ARTICLES } from "../store/actions.type";
export default {
  data() {
    return {
      offset: 0,
      limit: 10,
    };
  },
  computed: {
    ...mapState({
      articles: (state) => state.home.articles,
      articlesCount: (state) => state.home.articlesCount,
      type() {
        return this.$route.params.type
      }
    }),
  },
  method: {
    readDetail(slug) {
      console.log(slug)
      this.$route.push({name: "article", params:{slug: slug}})
    }
  },
  beforeMount() {
    this.$store.dispatch(FETCH_ARTICLES, {
      type: this.$route.params.type,
      offset: this.offset,
      limit: this.limit,
    })
  },
  watch: {
    type: function() {
      this.$store.dispatch(FETCH_ARTICLES, {
      type: this.$route.params.type,
      offset: this.offset,
      limit: this.limit,
    });
    }
  }
}
</script>

<style scoped>
  .info {
    margin-left: 16px;
  }

  .info .date {
    margin-top: 6px;
  }
</style>