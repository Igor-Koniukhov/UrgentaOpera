<template>
  <v-col lg="8" md="10">
    <validation-observer
        ref="observer"
        v-slot="{ invalid }"
    >
      <form @submit.prevent="submit">
        <validation-provider
            v-slot="{ errors }"
            name="Name"
            rules="required|max:10"
        >
          <v-text-field
              v-model="registration.name"
              :counter="10"
              :error-messages="errors"
              label="Name"
              required
          ></v-text-field>
        </validation-provider>
        <validation-provider
            v-slot="{ errors }"
            name="phoneNumber"
            :rules="{
          required: true,
          regex: '^[\\+]?[(]?[0-9]{3}[)]?[-\\s\\.]?[0-9]{3}[-\\s\\.]?[0-9]{4,6}$'
        }"
        >
          <v-text-field
              v-model="registration.phoneNumber"
              :counter="7"
              :error-messages="errors"
              label="Phone Number"
              required
          ></v-text-field>
        </validation-provider>
        <validation-provider
            v-slot="{ errors }"
            name="email"
            rules="required|email"
        >
          <v-text-field
              v-model="registration.email"
              :error-messages="errors"
              label="E-mail"
              required
          ></v-text-field>
        </validation-provider>

        <validation-provider
            v-slot="{ errors }"
            name="password"
            :rules="{
          required: true,
          regex: '^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$ %^&*-]).{8,}$'
        }"
        >
          <v-text-field
              v-model="registration.password"
              :error-messages="errors"
              label="Password"
              required
          ></v-text-field>
        </validation-provider>

        <validation-provider
            v-slot="{ errors }"
            name="confirmPassword"
        >
          <v-text-field
              v-model="registration.confirmPassword"
              :error-messages="errors"
              label="Password"
              required
          ></v-text-field>
        </validation-provider>

        <v-btn
            class="mr-4"
            type="submit"
            :disabled="invalid"
        >
          submit
        </v-btn>
        <v-btn @click="clear">
          clear
        </v-btn>
      </form>
    </validation-observer>
  </v-col>

</template>

<script>
import {digits, email, max, regex, required} from 'vee-validate/dist/rules'
import {extend, setInteractionMode, ValidationObserver, ValidationProvider} from 'vee-validate'

setInteractionMode('eager')

extend('digits', {
  ...digits,
  message: '{_field_} needs to be {length} digits. ({_value_})',
})

extend('required', {
  ...required,
  message: '{_field_} can not be empty',
})

extend('max', {
  ...max,
  message: '{_field_} may not be greater than {length} characters',
})

extend('regex', {
  ...regex,
  message: '{_field_} {_value_} does not match {regex}',
})

extend('email', {
  ...email,
  message: 'Email must be valid',
})
extend('password', {
  params: ['target'],
  validate(value, {target}) {
    return value === target;
  },
  message: 'Password confirmation does not match'
});

export default {
  components: {
    ValidationProvider,
    ValidationObserver,
  },
  data: () => ({
    registration:{
      name: '',
      phoneNumber: '',
      email: '',
      password: '',
      confirmPassword: ''
    },
    target: '',
    items: [
      'Item 1',
      'Item 2',
    ],

  }),

  methods: {
    submit() {
      this.$refs.observer.validate()
      this.$store.dispatch('auth/userRegistration', {
        name: this.registration.name,
        phoneNumber: this.registration.phoneNumber,
        email: this.registration.email,
        password: this.registration.password,
        confirmPassword: this.registration.confirmPassword,
      })
      this.clear()
      setTimeout(()=>{
        this.$router.push('/login')
      }, 2000)

    },
    clear() {
      this.registration.name = ''
      this.registration.phoneNumber = ''
      this.registration.email = ''
      this.registration.password = ''
      this.registration.confirmPassword = ''
      this.$refs.observer.reset()
    },
  },
}
</script>