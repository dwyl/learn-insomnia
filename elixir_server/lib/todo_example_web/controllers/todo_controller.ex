defmodule TodoExampleWeb.TodoController do
  use TodoExampleWeb, :controller

  def index(conn, _params) do
    json(conn, TodoExample.State.get_list)
  end

  def view(conn, %{"id" => id}) do
    json(conn, TodoExample.State.get_item(id))
  end

  def create(conn, params) do
    case Map.get(params, "item") do
      nil ->  conn |> put_status(400) |> json("Invalid parameters.")
      item -> json(conn, TodoExample.State.new_item(params))
    end
  end

  def admin(conn, _params) do
    token = List.first(get_req_header(conn, "authorization"))

    case token do
      "Bearer hunter2" -> text(conn, "Authorized!")
      _ -> text(conn, "Not authorized.")
    end
  end
end
