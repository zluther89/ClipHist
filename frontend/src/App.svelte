<script>
  let content = [];

  const fetchData = async () => {
    const res = await fetch("/content");
    content = await res.json();
  };

  const handleClick = async function(event, postInfo) {
    try {
      await fetch("/content", {
        method: "POST",
        body: JSON.stringify(postInfo)
      });
    } catch (e) {
      console.log(e);
    }
  };
  fetchData();
  setInterval(fetchData, 500);
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
      {#each content as d}
        <tr>
          <td class="date">{d.Timestamp}</td>
          <td id={d.Timestamp} on:click={e => handleClick(e, d)}>
            {d.Content}
          </td>
        </tr>
      {:else}
        <p>loading...</p>
      {/each}

    </tbody>
  </table>
</div>
