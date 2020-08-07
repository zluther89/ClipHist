<script>
  const fetchData = (async () => {
    const response = await fetch("/content");
    return await response.json();
  })();

  const handleClick = function(postInfo) {
    console.log(postInfo);
    try {
      fetch("/content", { method: "POST", body: JSON.stringify(postInfo) });
    } catch (e) {
      console.log(e);
    }
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
    background-color: rgb(78, 85, 88);
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
          <tr on:click={() => handleClick(d)}>
            <td class="date">{d.Timestamp}</td>
            <td>{d.Content}</td>
          </tr>
        {/each}
      {:catch error}
        {console.log(error)}
        <p>An error occurred!</p>
      {/await}

    </tbody>
  </table>
</div>
