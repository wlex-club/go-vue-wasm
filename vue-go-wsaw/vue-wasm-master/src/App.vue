<template>
  <div id="app">
    <p>{{num}}</p>
    <img src="./assets/logo.png" @click="test()">
    <HelloWorld :num="num"/>
  </div>
</template>

<script>
  import HelloWorld from './components/HelloWorld';

  export default {
    name: 'app',
    components: {
      HelloWorld,
    },
    data() {
      return {
        num: 0
      }
    },
    created() {
      if (this.$wasm.add) {
        console.log('zheli'); //
        console.log(this.$wasm.add.add(1, 2));
      }
    },
    methods: {
      test() {
        return this.num = this.$wasm.add.add(1, 2)
      }
    }
  };

</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }

</style>
