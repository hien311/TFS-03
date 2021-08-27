<template>
  <div id="order-summary" v-center ref="orderSummary">
      <img src="../assets/illustration-hero.svg" alt="illustration-hero">
      <div class="body-text">
          <h1>{{title}}</h1>
          <p>{{text}}</p>
      </div>
      <div class="plan" v-flex>
          <img src="../assets/icon-music.svg">
          <div class="plan-body">
              <h4>{{plan.name}}</h4>
              <p>{{plan.price}}</p>
          </div>
          <div class="plan-button" @click="changePlan">Change</div>
      </div>
      <div class="choice-buttons">
          <button @click="confirm">Proceed to Payment</button>
          <div class="cancel" @click="cancel">Cancel Order</div>
      </div>
      <Modal v-if="isConfirm" @close-modal="closeModal"/>
    </div>
</template>

<script>
import planData from "../mixins/planData"
export default {
    name: 'OderSummary',
    mixins: [planData],
    data() {
        return {
            title: "Order Summary",
            text: "You can now listen to millions of songs, audiobooks and podcasts on any device anywhere  you like!",
             plans: [
                {name: "Annual Plan", price:"$59.99/year"},
                {name: 'Pro Plan', price: '99.99/year'}
            ],
            planStyle: 0,
            isConfirm: false
        }
    },
    methods: {
        confirm() {
            this.isConfirm = true;
        },
        closeModal() {
            this.isConfirm = false
        },
        changePlan() {
            this.planStyle = ++this.planStyle % 2
        },
        cancel() {
            this.$refs['orderSummary'].style.display = 'none'
        },   
    },
    computed:{
        plan: function() {
            return this.plans[this.planStyle]
        }
    },
    directives: {
        flex: {
            bind: function(el) {
                el.style.display = 'flex'
            }
        },
        center: {
            inserted: function(el) {
                el.style['text-align'] = 'center'
            }
        },
        
    },
    components: { 
        Modal: () => import('./Modal.vue')
    },
}
</script>

<style scoped>
   #order-summary {
       width: 450px;
       height: 716px;
       margin: auto;
       border-radius: 18px;
       overflow: hidden;
       box-shadow: 5px 12px 12px rgba(207, 205, 205,1);
       background-color: white;
       
       position: fixed;
       top: 100px;
       left:50%;
       transform: translateX(-50%);
   }

    .body-text {
        margin-top: 54px;
        margin-bottom: 48px;
        line-height: 20px;
    }

    .body-text h1{
        font-weight: 700;
    }

    .body-text p {
        margin-top: 24px;
        letter-spacing: 0.3px;
        padding: 0 62px;
        opacity: 0.7;
    }

    .plan {
        height: 96px;
        width: 354px;

        justify-content: center;
        margin: default;
        padding: 24px 0;
        margin: auto;
        background: rgb(248,249,254);
        border-radius: 18px;
    }

    .plan img {
        height: 48px;
    }

    .plan .plan-body {
        margin-left: 20px;
        margin-right: 92px;
        text-align: start;
        line-height: 24px;
        min-width: 120px;
    }

    .plan-button {
        line-height: 48px;
        font-size: 16px;
        text-decoration: underline;
        color: rgb(63 53 179);
        font-weight:600;
    }

    .plan-button:hover {
        cursor: pointer;
        opacity: 0.9;
    }

    .choice-buttons {
        margin-top: 32px;
    }

    .choice-buttons button {
        width: 354px;
        height: 50px;
        border-radius: 12px;
        
        background-color: rgb(56,42,225);
        color: white;
        font-weight: 700;
        font-size: 16px;

        border: none;
        box-shadow: 4px 12px 12px rgba(207, 205, 205,1);
        
    }

    .choice-buttons button:hover {
        cursor: pointer;
        opacity: 0.9;
    }

    .choice-buttons .cancel {
        margin-top: 36px;
        color: rgb(159,165,191);
        font-weight: 600;
        font-size: 16px;
    }

    .choice-buttons .cancel:hover {
        cursor: pointer;
        opacity: 0.9;
    }
</style>