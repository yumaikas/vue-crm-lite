var app = new Vue({
    el: "#app",
    data: {
        message: 'Hello Vue!'
    },
});

var app2 = new Vue({
    el: "#app-2",
    data: {
        message: 'You loaded this page on ' + new Date(),
    },
});

var app3 = new Vue({
    el: '#app-3',
    data: {
        seen: true,
    }
});

var app4 = new Vue({
    el: '#app-4',
    data: {
        todos: [
            { text: 'Learn Javascript' },
            { text: 'Learn Vue' },
            { text: 'Build something awesome' },
        ],
    },
});

var app5 = new Vue({
    el: '#app-5',
    data: {
        message: 'Hello Vue.js!',
    },
    methods: {
        reverseMessage: function() {
            this.message = this.message.split('').reverse().join('');
        }
    }
});

var app6 = new Vue({
    el: '#app-6',
    data: {
        message: 'Hello Vue!',
    }
})

Vue.component('todo-item', {
    props: ['todo'],
    template: ' <li>{{ todo.text }} </li>'
});

var app7 = new Vue({
    el: '#app-7',
    data:{
        groceryList: [
            { text: 'Vegtables' },
            { text: 'Cheese' },
            { text: 'Whatever else humans are supposed to eat' },
        ],
    },
});

var appNums = (function() {
    function clamp(num, min, max) {
        if (num > max) { return max; }
        if (num < min) { return min; }
        return num;
    } 

    return new Vue({
        el: "#app-nums",
        data: {
            repetitions: 3
        },
        computed: {
            canAdd: function() {
                return this.repetitions < 10;
            },
            canRemove: function() {
                return this.repetitions > 1;
            },
        },
        methods: {
            addItem: function(){
                this.repetitions = clamp(this.repetitions + 1, 1, 10);
            },
            removeItem: function() {
                this.repetitions = clamp(this.repetitions -1, 1, 10);
            },
        },
    });
})();

