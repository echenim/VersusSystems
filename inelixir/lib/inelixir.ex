defmodule Inelixir do
  def take_home_assement do
    generator
    |> nameandpasswordtransformation
  end

  def generator do
    %Inelixir.Person{
      id: 123,
      name: "Elsa",
      username: "xXfrozen_princessXx",
      email: "elsa@arendelle.com",
      age: 21,
      power: "ice ray",
      friends: [
        %Inelixir.Friends{id: 234, username: "MagicSnowman32"},
        %Inelixir.Friends{id: 456, username: "call_me_anna"}
      ]
    }
  end

  def nameandpasswordtransformation(person) do
    %Inelixir.Person{
      id: id,
      name: name,
      username: username,
      email: email,
      age: age,
      power: power,
      friends: friends
    } = person

    %Inelixir.Person{
      id: id,
      name: name |> transformation,
      username: username |> transformation,
      email: email |> emailtranformation,
      age: age,
      power: power,
      friends: friends |> list_transformation
    }
  end

  def list_transformation(friends) do
    Enum.map(friends, fn x ->
      %Inelixir.Friends{
        id: x.id,
        username: x.username |> transformation
      }
    end)
  end

  def emailtranformation(email) do
    h =
      email
      |> String.split("@")
      |> hd
      |> transformation

    t = email |> String.split("@") |> tl
    "#{h}@#{t}"
  end

  def transformation(params) do
    Enum.map(String.graphemes(params), fn x -> "*" end) |> Enum.join("")
  end
end
