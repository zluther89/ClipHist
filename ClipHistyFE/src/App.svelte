<script>
  const fetchData = (async () => {
    const response = await fetch("/content");
    return await response.json();
  })();

  const handleClick = function(event, postInfo) {
    try {
      fetch("/content", { method: "POST", body: JSON.stringify(postInfo) });
    } catch (e) {
      console.log(e);
    }
  };

  var ws = new WebSocket("ws://" + window.location.host + "/socket");
  ws.onopen = function() {
    console.log("connected");
    ws.send(JSON.stringify({ message: "hello server!" }));
  };
  ws.onmessage = function(event) {
    var m = JSON.parse(event.data);
    console.log("Received message", m.message);
  };
  ws.onerror = function(event) {
    console.log(event);
  };
</script>

<style>
  table {
    margin-left: 10%;
    size: 50px;
    width: 60%;
    border-collapse: separate;
    border-spacing: 15px 20px;
  }
  td {
    padding-top: 1em;
    padding-bottom: 1em;
    padding-left: 1em;
    border-radius: 3px;
    box-shadow: rgba(15, 15, 15, 0.2) 0px 0px 0px 1px,
      rgba(15, 15, 15, 0.2) 0px 2px 4px;
    background-color: rgb(63, 68, 71);
  }
  th {
    font-size: 2em;
    text-align: left;
  }
  div {
    height: 100%;
    width: 100%;
    justify-content: center;
    align-items: center;
  }
  td:hover {
    background-color: #4e5558;
  }
  td:active {
    background-color: blue;
    /* background: linear-gradient(to middle, #4e5558 5%, #2f3437 100%); */
  }
  .date {
    width: 15%;
  }
</style>

<div>
  <table>
    <thead>
      <tr>
        <th class="date">date</th>
        <th>entry</th>
      </tr>
    </thead>
    <tbody>
      {#await fetchData}
        <p>...waiting</p>
      {:then data}
        {#each data as d}
          <tr>
            <td class="date">{d.Timestamp}</td>
            <td id={d.Timestamp} on:click={e => handleClick(e, d)}>
              {d.Content}
            </td>
          </tr>
        {/each}
      {:catch error}
        {console.log(error)}
        <p>An error occurred!</p>
      {/await}

    </tbody>
  </table>
</div>
