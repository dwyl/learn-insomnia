defmodule TodoExample.State do
  use Agent

  require Logger

  def start_link(_opts) do
    Agent.start_link(fn ->
      %{
        todos: [
          %{id: 1, item: "Star Dwyl on github"},
          %{id: 2, item: "Contribute to open source"}
        ],
        id_count: 3
      }
    end, name: __MODULE__)
  end

  def get_list() do
    Agent.get(__MODULE__, fn state -> state.todos end)
  end

  def get_item(id) do
    Agent.get(__MODULE__, fn state ->
      state
      |> Map.get(:todos)
      |> Enum.filter(fn item -> item[:id] == String.to_integer(id) end)
      |> List.first()
    end)
  end

  def get_id_count() do
    Agent.get(__MODULE__, &(&1.id_count))
  end

  def get_state() do
    Agent.get(__MODULE__, &(&1))
  end

  def new_item(%{"item" => item}) do
    new = %{
      id: get_id_count(),
      item: item
    }

    state = get_state()
    todos = state.todos

    state =
      state
      |> Map.replace!(:todos, [new | todos])
      |> Map.update!(:id_count, &(&1 + 1))

    Logger.debug(inspect state)
    Agent.update(__MODULE__, fn _ -> state end)

    new
  end
end
