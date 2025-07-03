components {
  id: "dungeon_guide"
  component: "/mobs/dungeon_guide/components/dungeon_guide.script"
}
components {
  id: "sprite"
  component: "/mobs/dungeon_guide/components/dungeon_guide.sprite"
}
components {
  id: "label"
  component: "/mobs/dungeon_guide/components/dungeon_guide.label"
  position {
    y: 8.0
  }
  scale {
    x: 0.0625
    y: 0.0625
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"mobs\"\n"
  "mask: \"mage\"\n"
  "mask: \"default\"\n"
  "mask: \"walls\"\n"
  "mask: \"mobs\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 20.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 7.9835\n"
  "  data: 27.2515\n"
  "  data: 10.2\n"
  "}\n"
  ""
}
embedded_components {
  id: "awareness"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"awareness\"\n"
  "mask: \"mage\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 39.443\n"
  "}\n"
  ""
}
