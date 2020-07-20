defmodule TodoExample.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  def start(_type, _args) do
    children = [
      # Start the Telemetry supervisor
      TodoExampleWeb.Telemetry,
      # Start the PubSub system
      {Phoenix.PubSub, name: TodoExample.PubSub},
      # Start the Endpoint (http/https)
      TodoExampleWeb.Endpoint,
      # Start a worker by calling: TodoExample.Worker.start_link(arg)
      # {TodoExample.Worker, arg}
      TodoExample.State
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: TodoExample.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  def config_change(changed, _new, removed) do
    TodoExampleWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
