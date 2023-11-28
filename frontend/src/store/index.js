import {createStore} from 'vuex'
import axios from 'axios'

export default createStore({
  state: () => ({
    allProducts: [
      {
        id: 1,
        image: [
          "src/assets/productPhotos/Nike/M2K Tekno/1.jpg",
          "src/assets/productPhotos/Nike/M2K Tekno/2.jpg",
          "src/assets/productPhotos/Nike/M2K Tekno/3.jpg",
          "src/assets/productPhotos/Nike/M2K Tekno/4.jpg"],
        brand: "Nike",
        model: "M2K Tekno",
        color: "white",
        price: "30.99$",
        description: "С новым взглядом на Nike Air Monarch мужские и женские кроссовки Nike M2K Tekno придают новое измерение любимым цветовым решениям в модной обуви"
      },
      {
        id: 2,
        image: [
          "src/assets/productPhotos/Nike/M2K Tekno/1.jpg",
          "src/assets/productPhotos/Nike/M2K Tekno/2.jpg",
          "src/assets/productPhotos/Nike/M2K Tekno/3.jpg",
          "src/assets/productPhotos/Nike/M2K Tekno/4.jpg"],
        brand: "Nike",
        model: "M2K Tekno",
        color: "white",
        price: "30.99$",
        description: "С новым взглядом на Nike Air Monarch мужские и женские кроссовки Nike M2K Tekno придают новое измерение любимым цветовым решениям в модной обуви"
      },
      {
        id: 3,
        image: [
          "src/assets/productPhotos/Nike/M2K Tekno/1.jpg",
          "src/assets/productPhotos/Nike/M2K Tekno/2.jpg",
          "src/assets/productPhotos/Nike/M2K Tekno/3.jpg",
          "src/assets/productPhotos/Nike/M2K Tekno/4.jpg"],
        brand: "Nike",
        model: "M2K Tekno",
        color: "white",
        price: "30.99$",
        description: "С новым взглядом на Nike Air Monarch мужские и женские кроссовки Nike M2K Tekno придают новое измерение любимым цветовым решениям в модной обуви"
      }
    ],
    cartProducts: [],
    user: {
      login: "",
      password: "",
      role: ""
    },
    drawer: false
  }),
  getters: {
    getProducts: (s) => s.allProducts,
  },
  mutations: {
    logInUser(logged){
      axios.get("")
      //if(userExist){
        this.user.login = logged.login
        this.user.password = logged.password
        this.user.role = logged.role
      //}
    },
    toggleDrawer(){
      this.drawer = !this.drawer
      console.log(this.drawer)
    }
  }
})


