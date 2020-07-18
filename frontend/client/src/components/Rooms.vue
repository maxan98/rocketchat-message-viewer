<template>

  <div class="container">
    <div class="row">
      <div class="col-sm-20">
        <h1>Books</h1>
        <hr><br><br>
        <button type="button" class="btn btn-success btn-sm">Add Book</button>
        <br><br>
        <table class="table table-hover ">
          <thead>
          <tr>
            <th scope="col">#</th>
            <th scope="col">Id</th>
            <th scope="col">Given Name</th>
            <th scope="col">Participants (if P2P)</th>
            <th></th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="(room, index) in rooms" :key="index">
            <td>{{ index }}</td>
            <td>{{ room.Id }}</td>
            <td>{{ room.Name }}</td>
            <td>
              <span>{{ room.Usernames }}</span>

            </td>
            <td>
              <button type="button" class="btn btn-success btn-sm " v-b-modal.room-modal @click="editBook(room)">Penetrate</button>


            </td>
          </tr>
          </tbody>
        </table>

      </div>
    </div>
    <b-modal ref="penetrateRoomModal"
             id="room-modal"
             title="Penentrate given room"
             size="huge"
             hide-footer>
      <b-form @submit="onSubmit" @reset="onReset" class="w-100">

              <table class="table table-hover ">
                <thead>
                <tr>

                  <th></th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(mess, index) in messages" :key="index">

                  <td>
                    <span>{{ mess.User }}</span>





                  </td>

                </tr>
                </tbody>
              </table>

        <q-btn label="My Button"></q-btn>

        <b-button type="close" variant="danger">Reset</b-button>

      </b-form>
    </b-modal>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    data() {
      return {
        rooms: [],
        messages: [],
        penetrateRoomForm: {
          title: '',
          author: '',
          read: [],
        },
      };
    },
    methods: {
      editBook(book) {
        const path = 'http://localhost:3000/api/v1/messages/'+book.Id;
        axios.get(path)
          .then((res) => {
            this.messages = res.data;
          })
          .catch((error) => {
            // eslint-отключение следующей строки
            console.error(error);
          });
        this.penetrateRoomForm.author = book.Id;
      },
      getRooms() {
        const path = 'http://localhost:3000/api/v1/rooms';
        axios.get(path)
          .then((res) => {
            this.rooms = res.data;
          })
          .catch((error) => {
            // eslint-отключение следующей строки
            console.error(error);
          });
      },
      ShowRooms(payload) {
        const path = 'http://localhost:5000/books';
        axios.post(path, payload)
          .then(() => {
            this.getRooms();
          })
          .catch((error) => {
            // eslint-отключение следующей строки
            console.log(error);
            this.getRooms();
          });
      },
      initForm() {
        this.penetrateRoomForm.title = '';
        this.penetrateRoomForm.author = '';
        this.penetrateRoomForm.read = [];
      },
      onSubmit(evt) {
        evt.preventDefault();
        this.$refs.penetrateRoomForm.hide();
        let read = false;
        if (this.penetrateRoomForm.read[0]) read = true;
        const payload = {
          title: this.penetrateRoomForm.title,
          author: this.penetrateRoomForm.author,
          read, // сокращённое свойство
        };
        this.ShowRooms(payload);
        this.initForm();
      },
      onReset(evt) {
        evt.preventDefault();
        this.$refs.penetrateRoomModal.hide();
        this.initForm();
      },
    },
    created() {
      this.getRooms();
    },
  };
</script>

<style>
  .modal .modal-huge {
    max-width: 1100px;
    width: 1100px;
  }
</style>

