defmodule InelixirTest do
  use ExUnit.Case
  doctest Inelixir

  test "greets the world" do
    assert Inelixir.hello() == :world
  end
end
