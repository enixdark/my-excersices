ExUnit.start()
import HelloWorld

defmodule HelloWorldTest do
  use ExUnit.Case, async: true
  # @tag :pending
  test "shouting" do
    assert HelloWorld.call "Hello, World!" == "Hello, World!"
  end
end
  