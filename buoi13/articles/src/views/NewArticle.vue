<template>
    <div id="new-article">
        <div class="editor-page">
            <div class="container page">
                <div class="row">
                    <div class="col-md-10 offset-md-1 col-xs-12">
                        <form>
                            <fieldset>
                                <fieldset class="form-group">
                                    <input
                                        type="text"
                                        class="form-control form-control-lg"
                                        placeholder="Article Title"
                                        v-model="title"
                                    />
                                </fieldset>
                                <fieldset class="form-group">
                                    <input
                                        type="text"
                                        class="form-control"
                                        placeholder="What's this article about?"
                                        v-model="description"
                                    />
                                </fieldset>
                                <fieldset class="form-group">
                                    <textarea
                                        class="form-control"
                                        rows="8"
                                        placeholder="Write your article (in markdown)"
                                        v-model="body"
                                    ></textarea>
                                </fieldset>
                                <fieldset class="form-group">
                                    <input
                                        type="text"
                                        class="form-control"
                                        placeholder="Enter tags"
                                        v-model="taglist"
                                    />
                                    <div class="tag-list"></div>
                                </fieldset>
                                <button
                                    class="btn btn-lg pull-xs-right btn-primary"
                                    type="button"
                                    @click="submit"
                                >
                                    Publish Article
                                </button>
                            </fieldset>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { POST_ARTICLES } from '../store/actions.type'
export default {
    data() {
        return {
            title: "",
            description: "",
            body: "",
            taglist: "",
        }
    },
    methods: {
        async submit() {
            this.taglist = this.taglist.split(", ")
            const isSuccess = await this.$store.dispatch(POST_ARTICLES, {
                title: this.title,
                description: this.description,
                body: this.body,
                taglist: this.taglist,
            })
            if (isSuccess) {
                let slug = this.$store.state.home.postResponse.article.slug
                this.$router.push({name: 'article', params: {slug}})
            }
        }
    }
}
</script>