<template>
  <div class="wrapper">
    <div class="brandRow">
      <v-btn-group floating elevation="2" class="rounded-b-xl" divided>
        <v-btn v-for="b in brandBtns" :key="b.id" class="brandBtn" :text="b.brand" />
      </v-btn-group>
    </div>
    <div class="sorters my-5">
      <v-row>
        <div class="sortField">
          <v-select label="Цвет" hide-details density="compact" :items="['Красный', 'Синий', 'Зеленый']"
            class="pr-2 sortField" />
        </div>
        <div class="sortField">
          <v-select label="Брэнд" hide-details density="compact" :items="['Adidas', 'Nike', 'Puma', 'Reebok']"
            class="sortField" />
        </div>
        <v-range-slider class="mx-5" :min="3000" :max="30000" thumb-label step="10" label="Цена"></v-range-slider>
        <div class="sortField">
          <v-select label="Сортировка" hide-details density="compact" :items="['Сначала дороже', 'Сначала дешевле']"
            class="sortField" />
        </div>
      </v-row>
    </div>
    <div class="products">
      <v-hover v-for="p in allProducts" :key="p.id">
        <template v-slot:default="{ isHovering, props }">
          <v-card :elevation="isHovering ? 5 : 0" v-bind="props" class="mx-auto my-2 rounded-lg" max-width="300px" >
            <v-img
              :src="p.image[0]"
              aspect-ratio="1/1" class="rounded-lg ma-2"/>
            <v-card-title :color="isHovering ? 'white' : 'grey'">{{p.brand}} {{p.model}}</v-card-title>
            <v-card-subtitle>
              {{ p.description }}
            </v-card-subtitle>
            <v-card-actions class="cardFooter">
              <v-btn variant="outlined" :color="!inCart ? '#354f52' : 'primary'"  @click="inCart = !inCart">{{!inCart ? "В корзину" : "Добавлено"}}</v-btn>
              <p>{{p.price}}</p>
            </v-card-actions>
          </v-card>
        </template>
      </v-hover>
    </div>
  </div>
</template>

<script>
import {mapMutations} from 'vuex'
import {mapGetters} from 'vuex'
export default {
  data: () => ({
    inCart: false,
    brandBtns: [
      {
        id: 1,
        brand: "adidas",
      },
      {
        id: 2,
        brand: "jordan",
      },
      {
        id: 3,
        brand: "nike",
      },
      {
        id: 4,
        brand: "puma",
      },
      {
        id: 5,
        brand: "reebok",
      },
      {
        id: 6,
        brand: "converse",
      },
      {
        id: 7,
        brand: "vans",
      },
      {
        id: 7,
        brand: "new balance",
      }
    ],
  }),
  methods: {
    ...mapMutations(['logInUser']),
  },
  computed: {
    ...mapGetters({
      allProducts: 'getProducts'
    }),
  }

}
</script>

<style scoped>
.brandRow {
  width: 100%;
  display: flex;
  justify-content: center;
}

.wrapper {
  width: 80%;
  align-items: center;
  margin-left: auto;
  margin-right: auto;
}

.sorters {
  width: 100%;
}

.sortField {
  width: 200px;
}

.products {

  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-around;
}

.productImage {
  border: solid black 1px;
  margin: 5px;
}
.cardFooter {
  display: flex;
  justify-content: space-between;
}
</style>
